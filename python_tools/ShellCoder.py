#!/usr/bin/python3
# -*- coding: utf-8 -*-
# a corrsponding tool can be found in the go_rat code framework... soonish?
# This file is being worked on but it shows how to use various methods to 
# generate shellcode. I will be expanding this to work in the go_rat framework
#
# Honestly, you should load the file into bpython and 
# instance the classes yourself
#   
################################################################################
##                      Makes Shellcode from Binary File                      ##
################################################################################                
# Licenced under GPLv3-modified                                               ##
# https://www.gnu.org/licenses/gpl-3.0.en.html                                ##
#                                                                             ##
# The above copyright notice and this permission notice shall be included in  ##
# all copies or substantial portions of the Software.                         ##
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
################################################################################
"""
Script to make shell code...

some encrytpion soonish , No salt

sincerely, mr_hai

"""

PROGRAM_DESCRIPTION = "Shellcoder"
TESTING = True

################################################################################
##############                    IMPORTS                      #################
################################################################################
import re
import r2pipe
import sys,os
import inspect
import traceback
import subprocess
import argparse
try:
    import colorama
    from colorama import init
    init()
    from colorama import Fore, Back, Style
# Not from the documentation on colorama
    if TESTING == True:
        COLORMEQUALIFIED = True
except ImportError as derp:
    herp_a = derp
    print("[-] NO COLOR PRINTING FUNCTIONS AVAILABLE, Install the Colorama Package from pip")
    COLORMEQUALIFIED = False

redprint          = lambda text: print(Fore.RED + ' ' +  text + ' ' + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
blueprint         = lambda text: print(Fore.BLUE + ' ' +  text + ' ' + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
greenprint        = lambda text: print(Fore.GREEN + ' ' +  text + ' ' + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
yellow_bold_print = lambda text: print(Fore.YELLOW + Style.BRIGHT + ' {} '.format(text) + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)

LOGLEVEL            = 'DEV_IS_DUMB'
LOGLEVELS           = [1,2,3,'DEV_IS_DUMB']

parser = argparse.ArgumentParser(description=PROGRAM_DESCRIPTION)
parser.add_argument('--file_input',
                                 dest    = 'FileInput',
                                 action  = "store" ,
                                 default = "./shellcode/reverse_connx64", 
                                 help    = "Binary file, currently only linux ELF supported " )
parser.add_argument('--disassembler',
                                 dest    = 'disassembler',
                                 action  = "store" ,
                                 default = "objdump", 
                                 help    = "Options can be objdump || radare2 , radare2 needs bpython to be useful currently" )
   
################################################################################
##############                 INTERNAL FUNkS                  #################
################################################################################
def error_printer(message):
    exc_type, exc_value, exc_tb = sys.exc_info()
    trace = traceback.TracebackException(exc_type, exc_value, exc_tb) 
    if LOGLEVEL == 'DEV_IS_DUMB':
        blueprint('LINE NUMBER >>>' + str(exc_tb.tb_lineno))
        greenprint('[+]The Error That Occured Was :')
        redprint( message + ''.join(trace.format_exception_only()))
        try:
            yellow_bold_print("Some info:")
            exc_info = sys.exc_info()
            traceback.print_exception(*exc_info)
            greenprint(traceback.format_tb(trace.exc_traceback))
            #greenprint(traceback.format_list(traceback.extract_tb(trace)[-1:])[-1])
        except Exception:
            yellow_bold_print("ERROR PRINTER FUCKED UP HERE IS WHY")
            redprint( message + ''.join(trace.format_exception_only()))
    else:
        redprint(message + ''.join(trace.format_exception_only()))

def exec_command(command, blocking = True, shell_env = True):
    '''Runs a command with subprocess.Popen'''
    try:
        if blocking == True:
            step = subprocess.Popen(command,shell=shell_env,stdout=subprocess.PIPE,stderr=subprocess.PIPE)
            output, error = step.communicate()
            for output_line in output.decode().split('\n'):
                print(output_line)
            for error_lines in error.decode().split('\n'):
                print(error_lines + " ERROR LINE")
        elif blocking == False:
            # TODO: not implemented yet                
            pass
    except Exception:
        error_printer("[-] Interpreter Message: exec_command() failed!")        

# metaclass to represent a disassembled file
class DisassembledFile():
    def __init__(self):
        pass
        

class Radare2Disassembler():
    def __init__(self, FileInput):
        herp = DisassembledFile()
        self.FileInput = FileInput
        self.radarpipe = r2pipe.open(FileInput)
        #self.data     = {}
        setattr(herp, "__name__", FileInput)
        setattr(herp, "Symbols", self.radarpipe.cmdj("isj"))
        setattr(herp, "Sections", self.radarpipe.cmdj("iSj"))
        setattr(herp, "Info", self.radarpipe.cmdj("ij"))
        setattr(herp, "arch", getattr(herp, "Info")["bin"]["arch"])
        setattr(herp, "bintype", getattr(herp, "Info")["bin"]["bintype"])
        setattr(herp, "bits", getattr(herp, "Info")["bin"]["bits"])
        setattr(herp, "binos", getattr(herp, "Info")["bin"]["os"])


class ObjDumpDisassembler():
    def __init__(self, FileInput):
        self.file_input = FileInput
        self.command = "objdump -d {} >> objdump-{}.txt".format(input,input)
        self.bestregexyet = "\t(?:[0-9a-f]{2} *){1,7}\t"
        self.start_of_main = "[0-9]*<main>:"
        self.hexstring = []
        self.asdf = []
        self.shellcode = []
        # do the thing
        self.exec_objdump(self.command)
        # get the hex
        try:
            self.objdump_input = open("objdump-{}.txt".format(self.file_input), "r")
        except Exception:
            error_printer("[-] Could not open file : objdump-{}.txt".format(self.file_input))
        try:
            self.ParseObjDumpOutput()
        except Exception:
            error_printer("[-] Could not parse ObjDump output")     
    def exec_objdump(self, input):
        ''' 
        Command to execute , place command line args here
        '''
        command = "objdump -d {} >> objdump-{}.txt".format(input,input)
        exec_command(command = command)
    def ParseObjDumpOutput(self):
        '''
        Returns a string with HEXCODES
        '''
        for line_of_text in self.objdump_input:
            if line_of_text != None:
                hex_match = re.search(self.bestregexyet,line_of_text, re.I)
                if hex_match != None:
                    self.shellcode.append(hex_match[0].replace("\t","").replace("\n", ""))
        for each in self.shellcode:
            self.asdf.append(each.strip())
        for line in self.asdf:
            for hexval in line.split(" "):
                self.hexstring.append("/x" + hexval)

class Disassembler():
    ''' main class that holds the logic for argparsing'''
    def __init__(self, filename, choice):
        if choice == "radare2":
            try:
                import r2pipe
                Radare2Disassembler(filename)
            except ImportError:
                error_printer("[-] R2PIPE not installed, falling back to objdump")
                ObjDumpDisassembler(filename)
        elif choice == "objdump":
            ObjDumpDisassembler(filename)
#        elif choice == "python":
#            PythonDisassembler(file_input=filename)

#class PythonDisassembler():
#    ''' pure python solution to getting binary as hex'''
#
#    def __init__(self, file_input):
#        try:
#            open(file= file_input)
#        except Exception:
#            error_printer("[-] Error: Failed to open file {}".format(file_input))
#class TestCompiler():
#    def __init__(self, input_src      = "cowtest.c",
#                       output_file    = "cowtest",
#                       compiler_flags = "-pthread"):
#        self.GCCCommand = 'gcc {} {} -o {}'.format(compiler_flags,
#                                                    input_src, 
#                                                    output_file)
#        #subprocess.Popen(self.GCCCommand)

#finding an executable to test on
#class DissTest():
#    def __init__(self):
#        pass

if __name__== "main":
    arguments = parser.parse_args()
    filename = arguments.FileInput
    choice = arguments.disassembler
    Disassembler(filename, choice)
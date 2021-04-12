#!/usr/bin/python3
# -*- coding: utf-8 -*-
# a corrsponding tool can be found in the go code framework
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
Script to make encrypted shell code...

No salt

sincerely, mr_hai

"""
TESTING = True

################################################################################
##############                    IMPORTS                      #################
################################################################################
import sys,os
import logging
import pkgutil
import inspect
import traceback
import subprocess
from pathlib import Path
from importlib import import_module
import argparse

PROGRAM_DESCRIPTION = "Shellcode"

parser = argparse.ArgumentParser(description=PROGRAM_DESCRIPTION)
parser.add_argument('--file_input',
                                 dest    = 'bin_file',
                                 action  = "store" ,
                                 default = "./a.out", 
                                 help    = "Binary file, currently only linux ELF supported " )
parser.add_argument('--salt',
                                 dest    = 'salt',
                                 action  = "store" ,
                                 default = "/dev/rand", 
                                 help    = "Salty jar of ... something, defaults to /dev/rand" )

parser.add_argument('--disassembler',
                                 dest    = 'disassembler',
                                 action  = "store" ,
                                 default = "radare2", 
                                 help    = "Options can be objdump or radare2" )

if __name__== "main":
    arguments = parser.parse_args
    
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

is_method          = lambda func: inspect.getmembers(func, predicate=inspect.ismethod)
LOGLEVEL            = 'DEV_IS_DUMB'
LOGLEVELS           = [1,2,3,'DEV_IS_DUMB']
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

class Disassembler():
    def __new__(cls, FileInput):
        if arguments.disassembler == "radare2":
            try:
                import r2pipe
            except ImportError:
                error_printer("[-] R2PIPE not installed, falling back to objdump")
                ObjDumpDisassembler(FileInput)
        elif arguments.disassembler == "objdump":
            ObjDumpDisassembler(FileInput)
    def __init__(self):
        pass

# metaclass to represent a disassembled file
class DisassembledFile():
    def __init__(self, FileInput):
        self.fileinput = FileInput
        


# dont use this yet, it's for another file
class Radare2Disassembler(Disassembler):
    def __init__(self, FileInput):
        self.FileInput = FileInput
        self.radarpipe = r2pipe.open(FileInput)
        self.data     = {}
        self.Symbols  = self.radarpipe.cmdj("isj")
        self.Sections = self.radarpipe.cmdj("iSj")
        self.Info     = self.radarpipe.cmdj("ij")
        self.arch     = self.Info["bin"]["arch"]
        self.bintype  = self.Info["bin"]["bintype"]
        self.bits     = self.Info["bin"]["bits"]
        self.binos    = self.Info["bin"]["os"]


class ObjDumpDisassembler(Disassembler):
    def __init__(self, file_input):
        self.file_input = file_input

    def exec_command(self, command, blocking = True, shell_env = True):
        '''Runs a command with subprocess.Popen'''
        try:
            if blocking == True:
                step = subprocess.Popen(command,shell=shell_env,stdout=subprocess.PIPE,stderr=subprocess.PIPE)
                output, error = step.communicate()
                for output_line in output.decode().split('\n'):
                    self.objdump_output = self.objdump_output + output_line
                for error_lines in error.decode().split('\n'):
                    self.objdump_error = self.objdump_error + error_lines
            elif blocking == False:
                # TODO: not implemented yet                
                pass
        except Exception:
            error_printer("[-] Interpreter Message: exec_command() failed!")        
    
    def exec_objdump(self, input):
        self.exec_command('objdump')


if __name__== "main":

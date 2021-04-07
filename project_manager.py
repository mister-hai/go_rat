#$/usr/bin/python3
# -*- coding: utf-8 -*-
################################################################################
##                      Makes project and builds binaries                     ##
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
Script to initialize a go project and build two target binaries
for a server/multiple-client infrastructure

This script makes Viruses, If you are using it to break the law, 

FUCK

YOU

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

try:
    import colorama
    from colorama import init
    init()
    from colorama import Fore, Back, Style
    if TESTING == True:
        COLORMEQUALIFIED = True
except ImportError as derp:
    print("[-] NO COLOR PRINTING FUNCTIONS AVAILABLE, Install the Colorama Package from pip")
    COLORMEQUALIFIED = False

################################################################################
##############                      VARS                       #################
################################################################################
PROJECT_NAME            = "go_rat"
PROJECT_DIRECTORY       = "/home/moop/Desktop/go_rat"
SHARED_CODE_DIRECTORY   = "/pkg/shared_code"
TARGET_SRC_DIRECTORY    = PROJECT_DIRECTORY + "/src/target_binary"
COMMAND_SRC_DIRECTORY   = PROJECT_DIRECTORY + "/src/command_binary"
# I want to modify some variables in the globals file 
# so we can compile custom binaries from generic code
# holding initial critical information
GLOBALS_FILE     = PROJECT_DIRECTORY + SHARED_CODE_DIRECTORY
possible_targets = {"windows": ["386","amd64","arm"],
                    "linux"  : ["386","amd64","arm","arm64"],
                    "android": ["386","amd64","arm","arm64"],
                    "darwin" : ["amd64","arm64"],
                    }
# TO BUILD FOR A DIFFERENT OS/ARCH
# CGO_ENABLED MUST BE SET TO "0"
CGOENV = 'CGO_ENABLED="1"'
# this should be set to the platform you want to build the binary for
BUILD_TARGET_OS      = "windows"
BUILD_TARGET_ARCH    = "amd64"
# set the folloiwing ENV vars to build specific targets
# otherwise GO compiler defaults to host specs
env_var_target_os    = os.environ["GOOS"]   = BUILD_TARGET_OS
env_var_target_arch  = os.environ["GOARCH"] = BUILD_TARGET_ARCH
# add entries as necessary to reflect go.mod file entries
PROJECT_DEPENDENCIES = ["github.com/fatih/color",
                        "github.com/hashicorp/mdns",
                        "golang.org/x/sys/windows/registry",
                        "github.com/shirou/gopsutil/process",
                        "github.com/cakturk/go-netstat/netstat",
                        "github.com/shirou/gopsutil/disk"]


LOGLEVEL            = 'DEV_IS_DUMB'
LOGLEVELS           = [1,2,3,'DEV_IS_DUMB']
log_file            = 'Go_rat_project'
logging.basicConfig(filename=log_file, format='%(asctime)s %(message)s', filemode='w')
logger              = logging.getLogger()
script_cwd          = Path().absolute()
script_osdir        = Path(__file__).parent.absolute()
################################################################################
##############               COLOR PRINTING                    #################
################################################################################

redprint          = lambda text: print(Fore.RED + ' ' +  text + ' ' + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
blueprint         = lambda text: print(Fore.BLUE + ' ' +  text + ' ' + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
greenprint        = lambda text: print(Fore.GREEN + ' ' +  text + ' ' + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
yellow_bold_print = lambda text: print(Fore.YELLOW + Style.BRIGHT + ' {} '.format(text) + Style.RESET_ALL) if (COLORMEQUALIFIED == True) else print(text)
makeyellow        = lambda text: Fore.YELLOW + ' ' +  text + ' ' + Style.RESET_ALL if (COLORMEQUALIFIED == True) else text
makered           = lambda text: Fore.RED + ' ' +  text + ' ' + Style.RESET_ALL if (COLORMEQUALIFIED == True) else None
makegreen         = lambda text: Fore.GREEN + ' ' +  text + ' ' + Style.RESET_ALL if (COLORMEQUALIFIED == True) else None
makeblue          = lambda text: Fore.BLUE + ' ' +  text + ' ' + Style.RESET_ALL if (COLORMEQUALIFIED == True) else None
debug_message     = lambda message: logger.debug(blueprint(message)) 
info_message      = lambda message: logger.info(greenprint(message))   
warning_message   = lambda message: logger.warning(yellow_bold_print(message)) 
error_message     = lambda message: logger.error(redprint(message)) 
critical_message  = lambda message: logger.critical(yellow_bold_print(message))

is_method          = lambda func: inspect.getmembers(func, predicate=inspect.ismethod)

################################################################################
##############                 INTERNAL FUNCS                  #################
################################################################################
def error_printer(message):
    exc_type, exc_value, exc_tb = sys.exc_info()
    trace = traceback.TracebackException(exc_type, exc_value, exc_tb) 
    if LOGLEVEL == 'DEV_IS_DUMB':
        debug_message('LINE NUMBER >>>' + str(exc_tb.tb_lineno))
        greenprint('[+]The Error That Occured Was :')
        error_message( message + ''.join(trace.format_exception_only()))
        try:
            critical_message("Some info:")
            exc_info = sys.exc_info()
            traceback.print_exception(*exc_info)
            makegreen(traceback.format_tb(trace.exc_traceback))
            #makegreen(traceback.format_list(traceback.extract_tb(trace)[-1:])[-1])
        except Exception:
            critical_message("ERROR PRINTER FUCKED UP HERE IS WHY")
            error_message( message + ''.join(trace.format_exception_only()))
    else:
        error_message(message + ''.join(trace.format_exception_only()))

def exec_command(command, blocking = True, shell_env = True):
    '''Runs a command with subprocess.Popen'''
    try:
        if blocking == True:
            step = subprocess.Popen(command,shell=shell_env,stdout=subprocess.PIPE,stderr=subprocess.PIPE)
            output, error = step.communicate()
            for output_line in output.decode().split('\n'):
                info_message(output_line)
            for error_lines in error.decode().split('\n'):
                critical_message(error_lines)
        elif blocking == False:
            # TODO: not implemented yet                
            pass
    except Exception:
        error_printer("[-] Interpreter Message: exec_command() failed!")

################################################################################
##############                 MEAT N TATERS                   #################
################################################################################
# PYTHON3 script to initialize a new project with the go_rat module
class GoRatManager():
    def __init__(self):
        self. start_menu = '''
Enter your selection using a single integer:

1. Initialize project : {project_name}
2. Install Dependencies to GOPATH
3. Build Target for Windows
4. Build Command And Control For Linux
'''.format(project_name = PROJECT_NAME)
        blueprint(self.start_menu)
        menu_selection = input()
        # yes I validate like a fool
        # users are foolish
        if str.isdigit(menu_selection)     \
          and (len(menu_selection) == 1) \
          and menu_selection > 4:
            if menu_selection ==1:
                pass
            elif menu_selection == 2:
                pass
            elif menu_selection == 3:
                pass
            elif menu_selection == 4:
                pass
    
    def edit_globals(self, globals_file = GLOBALS_FILE):
        file_to_modify = open(globals_file, "w")

    def init_project(self):
        '''Initializes the folder this script resides in as a go project'''
        os.chdir(PROJECT_DIRECTORY)
        subprocess.Popen("go mod init {}".format(PROJECT_NAME))

    def install_dependencies(self, utility_to_use = "go get"):
        if utility_to_use == "go get":
            for dependency_url in PROJECT_DEPENDENCIES:
                exec_command("go get {}".format(dependency_url))

    def build_zombie_for_target(self, target_arch: str, target_os : str):
        '''fed with values from the variables at the top of this file '''
        os.chdir(TARGET_SRC_DIRECTORY)
        exec_command("go build {}".format())
        pass
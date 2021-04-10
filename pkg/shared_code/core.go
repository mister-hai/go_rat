/*/
This file contains the functions necessary for new commands and stuff lol

/*/
package shared_code

/*/
Must have a way of constructing new structs,
they are structures, you must build them
One for each in this project
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

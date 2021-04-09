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
func NewHostIntel() *HostIntel {
	new_host_intel := HostIntel{}
	return &new_host_intel
}
func NewOSInfo() *OSInfo {
	new_os_info := OSInfo{}
	return &new_os_info
}

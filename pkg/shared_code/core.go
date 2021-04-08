/*/
This file contains the functions necessary for new commands and stuff lol

/*/
package shared_code

import "encoding/json"

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
	new_command.json_input = json.RawMessage(json_input)

}

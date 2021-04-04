/*/
This file contains the global variables we need to allow our functions to
share various things between them. Try to avoid having too many of these.
And limit thier scope/context within which you use them
/*/
package go_rat

// import the libraries we need
import (
	"crypto"
	"net"

	"github.com/hashicorp/mdns"
	/*/  IMPORTING MODULES YOU FIND ONLINE
		in the terminal in VSCODE, while in the package root directory,
		append the following imports, as is, to the command "go get"

		Example:

		go get github.com/hashicorp/mdns

		And it will install the modules to the
		GOMODCACHE environment variable

	/*/// for colorized printing
	// basic ANSI Escape sequences
	// necessary for multicast DNS
)

// declaring global variables to share our
// network information between scopes
// these are for TCP/UDP specifically
// instanced without a value assigned
var local_tcpaddr_LAN net.TCPAddr
var local_tcpaddr_WAN net.TCPAddr
var local_udpaddr_LAN net.UDPAddr
var local_udpaddr_WAN net.UDPAddr

// Command And Control
// At the top level scope (module level)
// you declare with a simple "="

// instanced with a value assigned
var remote_tcpport string = ":1337"
var remote_tcpaddr string = "192.168.0.2" + remote_tcpport
var remote_udpport string = ":1337"
var remote_udpaddr string = remote_tcpaddr + remote_udpport
var PHONEHOME_TCP net.TCPAddr
var PHONEHOME_UDP net.UDPAddr
var mdns_server mdns.Server

//-----NAME-------------TYPE-----

// Admin Password in an obvious place
// TODO: set these for "hardmode" section
var sha256_admin_pass_preencrypted crypto.Hash
var sha512_admin_pass_preencrypted crypto.Hash

// Horribly insecure implementation
var sha256_hash_admin crypto.Hash
var new_admin_hash = sha256_hash_admin.New()
var wat = new_admin_hash.Sum([]byte("admin"))

/******************************************************************************
 This is a Remote Administration Tool, Command And Control Binary

    written in Golang

    As a practice in learning Golang

    And as a tutorial for "the Church Of The Subhacker" Wiki

    This tutorial assumes some familiarity with programming concepts,
	languages, and networking

	This is the other half of a two part tutorial
	on programming for hackers... just as soon as I finish lol

=================================================================
/*/

// we have to name the module after the folder its in
package command_binary

// import the libraries we need

/*/  IMPORTING MODULES YOU FIND ONLINE
	in the terminal in VSCODE, while in the package root directory,
	append the following imports, as is, to the command "go get"

	Example:

	go get github.com/hashicorp/mdns

	And it will install the modules to the
	GOMODCACHE environment variable

/*/ // for colorized printing
// basic ANSI Escape sequences
// necessary for multicast DNS

/*/ declaring global variables to share our
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
/*/

// we need objects to represent a compromised computer
type Zombie struct {
}

// and a set of compromised computers
type ZombieHorde struct {
}

// and a function to perform actinos with them
func something_nice() {

}

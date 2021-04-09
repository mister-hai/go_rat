// +build linux

/*/
This is the file that will be compiled into the binary
	that gets placed on the target host

This file contains the code only used for building
	- LINUX BINARIES

We are using the feature outlined in :
	- https://golang.org/cmd/go/#hdr-Build_constraints

	Placing the tag "+build some_tag_here"
		as a comment "//"
		at the top of the file
		and using "go build -tags some_tag_here"

	will compile main.go with code from that file mixed in

/*/
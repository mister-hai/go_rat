#!/bin/bash

#The build script will build a generic command and control executable file
# and a custom zombie executable, that has various values set
# BEACON_ON_START  bool
# BACON_TYPE       string ("tcp","http","udp","dns")

# to build the binaries for each of the two packages, without installing 
# them to $GOPATH/bin, then you will need to cd in to the paths where the 
# main.go files are, and run go build there:
cd target_binary

# now you need to build the binary that goes on the exploited host, assuming it is a windows machine,
# You may have chosen something different for your setup
# GOOS    : operating system (Linux, Windows, BSD, etc.)
# GOARCH  : architecture to build for.
#   - By default "go build" will generate an executable for the current platform and architecture. 
#   - This is being done on linux VM, so if you want a win64 exe, you need to do the following;
env GOOS=linux GOARCH=arm64 go build -o prepnode_arm64
go build main.go


# now make the command and control binary, there will be several errors you must fix:
# some of these errors are pointed out with a comment, some are not.


#$/bin/bash


# to build the binaries for each of the two packages, without installing 
# them to $GOPATH/bin, then you will need to cd in to the paths where the 
# main.go files are, and run go build there:

# To build the binary that goes on the exploited host:
cd target_binary;
go build main.go

# now make the command and control binary, there will be several errors you must fix:
# some of these errors are pointed out with a comment, some are not.
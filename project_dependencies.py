#!/usr/bin/python3
# this file gets called by "importmodule"

# this file is the container for the list of project dependencies
# so you dont have to fuck around with getting things working.
GO_DEPENDENCIES = ["github.com/fatih/color",
                        "github.com/hashicorp/mdns",
                        "golang.org/x/sys/windows/registry",
                        "github.com/shirou/gopsutil/process",
                        "github.com/cakturk/go-netstat/netstat",
                        "github.com/shirou/gopsutil/disk",
                        "github.com/godbus/dbus/v5",
                        "github.com/rainycape/dl",
                        "github.com/sirupsen/logrus"]

#installed via "pip3 install"
PYTHON_DEPENDENCIES = ["donut-shellcode","","","",""]
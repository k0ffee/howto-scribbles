package main

// keepalived-state -- simply write keepalived's virtual sync group or
// interface status into files for later processing by monitoring systems.
//
// Call from keepalive with "notify", like:
//   virtual_interface example {
//       notify "/usr/local/bin/keepalived-state"
//   }
//
// This would write to file "/var/run/keepalived-state.INTERFACE.example

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var dir string = "/var/run"
var file string = "keepalived-state"

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s arg1 arg2 arg3\n", os.Args[0])
		os.Exit(1)
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]

	data := []byte(fmt.Sprintln(arg1, arg2, arg3))

	name := path.Clean(fmt.Sprintf("%s/%s.%s.%s", dir, file, arg1, arg2))

	err := ioutil.WriteFile(name, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", name)
		os.Exit(1)
	}
}

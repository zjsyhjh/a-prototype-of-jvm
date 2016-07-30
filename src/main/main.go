package main

import (
	"fmt"
)

func startJVM(cmd *Cmd) {
	fmt.Println("JVM is starting...")
}

/*
 *
 */
func main() {

	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.1")
	} else if cmd.helpFlag || cmd.className == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

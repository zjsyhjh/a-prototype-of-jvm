package main

import (
	"fmt"
	"jvm/classpath"
	"strings"
)

func startJVM(cmd *Cmd) {
	fmt.Println("JVM is starting...")
	cp := classpath.ParseClassPathOption(cmd.xjreOption, cmd.cpOptinon)
	fmt.Printf("%v\n", cp)
	fmt.Printf("class : %v, args : %v\n", cmd.className, cmd.classArgs)

	className := strings.Replace(cmd.className, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Coundn't find or load the Main class : %v\n", cmd.className)
		return
	}

	fmt.Println("SUCCEED")
	fmt.Printf("Class data : %v\n", classData)
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

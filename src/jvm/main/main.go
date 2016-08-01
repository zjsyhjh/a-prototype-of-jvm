package main

import (
	"fmt"
	"jvm/classpath"
	"strings"
)

func startJVM(cmd *Cmd) {
	fmt.Println("JVM is starting...")
	cp := classpath.ParseClassPathOption(cmd.XjreOption, cmd.cpOptinon)
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
 * 执行go install /path/dir/*
 * 在bin目录下，可以看到可执行文件main, 执行./main [options] class [args...]
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

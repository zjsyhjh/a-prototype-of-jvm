package main

import (
	"fmt"
	"jvm/classpath"
	"jvm/rtda/heap"
	"strings"
)

/*
func parseClassFile(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Coundn't find or load the Main class, %v\n", className)
		panic(err)
	}

	cf, err := classfile.ParseClassFile(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassFileInfo(cf *classfile.ClassFile) {
	fmt.Printf("Version : %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("Constants count : %v\n", len(cf.ConstantPool().ConstantPool()))
	fmt.Printf("Access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("This class : %v\n", cf.ClassName())
	fmt.Printf("Super class : %v\n", cf.SuperClassName())
	fmt.Printf("Interfaces : %v\n", cf.InterfaceNames())
	fmt.Printf("Fields count : %v\n", len(cf.Fields()))
	for _, field := range cf.Fields() {
		fmt.Printf("%s\n", field.Name())
	}
	fmt.Printf("Methods count : %v\n", len(cf.Methods()))
	for _, method := range cf.Methods() {
		fmt.Printf("%s\n", method.Name())
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, method := range cf.Methods() {
		if method.Name() == "main" && method.Descriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
}

*/

func startJVM(cmd *Cmd) {
	//fmt.Println("JVM is starting...")

	cp := classpath.ParseClassPathOption(cmd.XjreOption, cmd.cpOptinon)
	/*
		fmt.Printf("%v\n", cp)
		fmt.Printf("class : %v, args : %v\n", cmd.className, cmd.classArgs)
	*/
	className := strings.Replace(cmd.className, ".", "/", -1)
	/*
		cf := parseClassFile(className, cp)
		printClassFileInfo(cf)
		mainMethod := getMainMethod(cf)
	*/
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod == nil {
		fmt.Printf("Main method couldn't found in class %s\n", cmd.className)
	} else {
		interpret(mainMethod, cmd.verboseInstFlag, cmd.classArgs)
	}
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

package main

import (
	"flag"
	"fmt"
	"os"
)

/*
 * 命令行参数选项
 * -Xjre选项输入时不能用通用符
 */
type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	className        string
	classArgs        []string
	cpOptinon        string
	XjreOption       string
}

/*
 * 一般java命令行格式：java [-options] class [args...]
 * verboseClassFlag用于控制是否把类加载信息输出到控制台
 * verboseInstFlag用于控制是否把指令执行信息输出到控制台
 *
 */

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "this is help information")
	flag.BoolVar(&cmd.helpFlag, "?", false, "this is help information")
	flag.BoolVar(&cmd.versionFlag, "version", false, "this is version information")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOptinon, "classpath", "", "the user's classpath")
	flag.StringVar(&cmd.cpOptinon, "cp", "", "the user's classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	/*
	 * 必须放在标志性参数后面，否则会出现解析错误
	 * 标志性参数如 -cp="xxx", 非标志性参数如 arg1, arg2
	 */
	flag.Parse()

	/*
	 * 取得非标志参数
	 */
	args := flag.Args()

	if len(args) > 0 {
		cmd.className = args[0]
		cmd.classArgs = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	fmt.Println("-version : show the version of jvm")
	fmt.Println("-classpath: use this -classpath=classpath")
	fmt.Println("-cp : use this -cp=classpath")
	fmt.Println("-verbose:class : use this -verbose:class to control output")
}

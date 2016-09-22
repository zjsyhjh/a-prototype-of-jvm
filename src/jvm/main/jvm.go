package main

import (
	"fmt"
	"jvm/classpath"
	"jvm/rtda"
	"jvm/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.ParseClassPathOption(cmd.XjreOption, cmd.cpOptinon)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (self *JVM) start() {
	/*
		self.initVM()
	*/
	self.execMain()
}

/*
func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}
*/

func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.className, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod == nil {
		fmt.Println("Main method not found in class %s\n", self.cmd.className)
		return
	}

	argsArray := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArray)

	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

/*
 * 把参数转换成java字符串数组
 */
func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.classArgs))
	argsArray := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArray.Refs()
	for i, arg := range self.cmd.classArgs {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArray
}

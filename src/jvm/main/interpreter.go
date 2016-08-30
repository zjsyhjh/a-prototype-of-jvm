package main

import (
	"fmt"
	"jvm/instructions"
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		/*
			fmt.Printf("LocalVars : %v\n", frame.LocalVars())
			fmt.Printf("OperandStack : %v\n", frame.OperandStack())
			panic(r)
		*/
		logFrames(thread)
		panic(r)
	}
}

/*
 * loop函数的循环执行: 计算PC， 解码指令， 执行指令
 */
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		/*
		 * decode
		 */
		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}
		/*
		 * execute
		 */
		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

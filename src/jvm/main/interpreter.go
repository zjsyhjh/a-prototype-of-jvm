package main

import (
	"fmt"
	"jvm/classfile"
	"jvm/instructions"
	"jvm/instructions/base"
	"jvm/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, byteCode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars : %v\n", frame.LocalVars())
		fmt.Printf("OperandStack : %v\n", frame.OperandStack())
		panic(r)
	}
}

/*
 * loop函数的循环执行: 计算PC， 解码指令， 执行指令
 */
func loop(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		/*
		 * decode
		 */
		reader.Reset(byteCode, pc)
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		/*
		 * execute
		 */
		fmt.Printf("pc : %2d, inst : %T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

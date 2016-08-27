package controls

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * return指令，从虚拟机栈中弹出
 */
type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

/*
 * 返回reference
 */
type ARETURN struct {
	base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currrentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	ref := currrentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(ref)
}

/*
 * 返回int类型
 */
type IRETURN struct {
	base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	value := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(value)
}

/*
 * 返回long类型
 */
type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	value := currentFrame.OperandStack().PopLong()
	invokeFrame.OperandStack().PushLong(value)
}

/*
 * 返回float类型
 */
type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	value := currentFrame.OperandStack().PopFloat()
	invokeFrame.OperandStack().PushFloat(value)
}

/*
 * 返回double类型
 */
type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	value := currentFrame.OperandStack().PopDouble()
	invokeFrame.OperandStack().PushDouble(value)
}

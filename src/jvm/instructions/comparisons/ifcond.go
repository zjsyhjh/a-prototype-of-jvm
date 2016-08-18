package comparisons

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * if<cond>指令把操作数从栈顶弹出，然后跟0比较，满足条件则跳转
 * ifeq == 0
 * ifne != 0
 * iflt < 0
 * ifle <= 0
 * ifgt > 0
 * ifge >= 0
 */
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	value := frame.OperandStack().PopInt()
	if value == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	value := frame.OperandStack().PopInt()
	if value != 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	value := frame.OperandStack().PopInt()
	if value < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	value := frame.OperandStack().PopInt()
	if value <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	value := frame.OperandStack().PopInt()
	if value > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	value := frame.OperandStack().PopInt()
	if value >= 0 {
		base.Branch(frame, self.Offset)
	}
}

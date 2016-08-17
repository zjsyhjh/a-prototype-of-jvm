package stacks

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * pop和pop2指令将栈顶变量弹出
 * pop指令只能用于弹出int、float等占用一个操作数栈位置的变量
 * long, double变量在操作数栈中占据两个位置，需要使用pop2指令弹出
 */
type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	s.PopSlot()
	s.PopSlot()
}

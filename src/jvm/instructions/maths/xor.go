package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 将栈顶的两个int型数值「按位异或」之后将结果入栈
 */
type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopInt()
	var2 := s.PopInt()
	s.PushInt(var1 ^ var2)
}

/*
 * 将栈顶的两个long型数值「按位异或」之后将结果入栈
 */
type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopLong()
	var2 := s.PopLong()
	s.PushLong(var1 ^ var2)
}

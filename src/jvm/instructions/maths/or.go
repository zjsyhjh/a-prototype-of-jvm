package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 将栈顶的两个int型数值「按位或」之后将结果入栈
 */
type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopInt()
	var2 := s.PopInt()
	s.PushInt(var1 | var2)
}

/*
 * 将栈顶的两个long型数值「按位或」之后将结果入栈
 */
type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopLong()
	var2 := s.PopLong()
	s.PushLong(var1 | var2)
}

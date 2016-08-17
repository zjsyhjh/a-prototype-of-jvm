package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 将栈顶的两个int型数值「按位与」之后将结果入栈
 */
type IAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopInt()
	var2 := s.PopInt()
	s.PushInt(var1 & var2)
}

/*
 * 将栈顶的两个long型数值「按位与」之后将结果入栈
 */
type LAND struct {
	base.NoOperandsInstruction
}

func (self *LAND) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopLong()
	var2 := s.PopLong()
	s.PushLong(var1 & var2)
}

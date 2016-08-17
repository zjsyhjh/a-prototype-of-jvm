package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * int类型相加
 */
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopInt()
	var2 := s.PopInt()
	s.PushInt(var1 + var2)
}

/*
 * long类型相加
 */
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopLong()
	var2 := s.PopLong()
	s.PushLong(var1 + var2)
}

/*
 * float类型相加
 */
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopFloat()
	var2 := s.PopFloat()
	s.PushFloat(var1 + var2)
}

/*
 * double类型相加
 */
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopDouble()
	var2 := s.PopDouble()
	s.PushDouble(var1 + var2)
}

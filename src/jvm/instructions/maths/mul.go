package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * int型相乘
 */
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopInt()
	var2 := s.PopInt()
	s.PushInt(var1 * var2)
}

/*
 * long型相乘
 */
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopLong()
	var2 := s.PopLong()
	s.PushLong(var1 * var2)
}

/*
 * float型相乘
 */
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopFloat()
	var2 := s.PopFloat()
	s.PushFloat(var1 * var2)
}

/*
 * double型相乘
 */
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var1 := s.PopDouble()
	var2 := s.PopDouble()
	s.PushDouble(var1 * var2)
}

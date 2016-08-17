package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * int型相减
 */
type ISUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopInt()
	s.PushInt(var1 - var2)
}

/*
 * long型相减
 */
type LSUB struct {
	base.NoOperandsInstruction
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopLong()
	var1 := s.PopLong()
	s.PushLong(var1 - var2)
}

/*
 * float型相减
 */
type FSUB struct {
	base.NoOperandsInstruction
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopDouble()
	var1 := s.PopDouble()
	s.PushDouble(var1 - var2)
}

/*
 * double型相减
 */
type DSUB struct {
	base.NoOperandsInstruction
}

func (self *DSUB) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopDouble()
	var1 := s.PopDouble()
	s.PushDouble(var1 - var2)
}

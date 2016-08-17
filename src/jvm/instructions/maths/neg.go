package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * int型取反
 */
type INEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	val := s.PopInt()
	s.PushInt(-val)
}

/*
 * long型取反
 */
type LNEG struct {
	base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	val := s.PopLong()
	s.PushLong(-val)
}

/*
 * float型取反
 */
type FNEG struct {
	base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	val := s.PopFloat()
	s.PushFloat(-val)
}

/*
 * double型取反
 */
type DNEG struct {
	base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	val := s.PopDouble()
	s.PushDouble(-val)
}

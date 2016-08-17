package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * int型相除
 */
type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopInt()
	if var2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	s.PushInt(var1 / var2)
}

/*
 * long型相除
 */
type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopLong()
	var1 := s.PopLong()
	if var2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	s.PushLong(var1 / var2)
}

/*
 * float型相除
 */
type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopFloat()
	var1 := s.PopFloat()
	s.PushFloat(var1 / var2)
}

/*
 * double型相除
 */
type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopDouble()
	var1 := s.PopDouble()
	s.PushDouble(var1 / var2)
}

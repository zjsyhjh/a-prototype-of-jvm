package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"math"
)

/*
 * int型取余
 */
type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopInt()
	if var2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	s.PushInt(var1 % var2)
}

/*
 * long型取余
 */
type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopLong()
	var1 := s.PopLong()
	if var2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	s.PushLong(var1 % var2)
}

/*
 * float型取余
 */
type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopFloat()
	var1 := s.PopFloat()
	res := float32(math.Mod(float64(var1), float64(var2)))
	s.PushFloat(res)
}

/*
 * double型取余
 */
type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopDouble()
	var1 := s.PopDouble()
	s.PushDouble(math.Mod(var1, var2))
}

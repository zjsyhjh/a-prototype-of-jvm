package conversions

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * long -> int
 */
type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopLong()
	s.PushInt(int32(value))
}

/*
 * long -> float32
 */
type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopLong()
	s.PushFloat(float32(value))
}

/*
 * long -> double
 */
type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopLong()
	s.PushDouble(float64(value))
}

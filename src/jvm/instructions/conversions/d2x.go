package conversions

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * double -> int
 */
type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopDouble()
	s.PushInt(int32(value))
}

/*
 * double -> long
 */
type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopDouble()
	s.PushLong(int64(value))
}

/*
 * double -> float
 */
type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopDouble()
	s.PushFloat(float32(value))
}

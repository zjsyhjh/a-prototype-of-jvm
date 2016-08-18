package conversions

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * float -> int
 */
type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2I) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopFloat()
	s.PushInt(int32(value))
}

/*
 * float -> long
 */
type F2L struct {
	base.NoOperandsInstruction
}

func (self *F2L) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopFloat()
	s.PushLong(int64(value))
}

/*
 * float -> double
 */
type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	value := s.PopFloat()
	s.PushDouble(float64(value))
}

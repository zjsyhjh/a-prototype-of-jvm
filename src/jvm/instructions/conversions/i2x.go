package conversions

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 强制类型转换，int -> 其他类型
 */
/*
 * int -> byte
 */
type I2B struct {
	base.NoOperandsInstruction
}

func (self *I2B) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	i := s.PopInt()
	s.PushInt(int32(int8(i)))
}

/*
 * int -> char
 */
type I2C struct {
	base.NoOperandsInstruction
}

func (self *I2C) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	i := s.PopInt()
	s.PushInt(int32(uint16(i)))
}

/*
 * int -> short
 */
type I2S struct {
	base.NoOperandsInstruction
}

func (self *I2S) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	i := s.PopInt()
	s.PushInt(int32(int16(i)))
}

/*
 * int -> long
 */
type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	i := s.PopInt()
	s.PushLong(int64(i))
}

/*
 * int -> float
 */
type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	i := s.PopInt()
	s.PushFloat(float32(i))
}

/*
 * int -> double
 */
type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	i := s.PopInt()
	s.PushDouble(float64(i))
}

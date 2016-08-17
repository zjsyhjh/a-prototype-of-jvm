package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * var1是要进行位移操作的变量，var2是指出位移多少比特, 位移之后的结果入栈
 * int变量只有32位，所以只取var2的前5个比特，long变量64位，取var2的前6个比特
 * go语言位移操作符右侧必须为无符号整数，因此对var2要进行类型转换
 * go语言中没有>>>逻辑右移运算符，因此为达到无符号目的，需要先把var1转换成无符号整数，位移操作之后，再转回有符号整数
 */

/*
 * int型算数左移
 */
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopInt()
	bits := uint32(var2) & 0x1f
	s.PushInt(var1 << bits)
}

/*
 * int型算数右移
 */
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopInt()
	bits := uint32(var2) & 0x1f
	s.PushInt(var1 >> bits)
}

/*
 * int型逻辑右移
 */
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopInt()
	bits := uint32(var2) & 0x1f
	s.PushInt(int32(uint32(var1) >> bits))
}

/*
 * long型算数左移
 */
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopLong()
	bits := uint32(var2) & 0x3f
	s.PushLong(var1 << bits)
}

/*
 * long型算数右移
 */
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopLong()
	bits := uint32(var2) & 0x3f
	s.PushLong(var1 >> bits)
}

/*
 * long型逻辑右移
 */
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopInt()
	var1 := s.PopLong()
	bits := uint32(var2) & 0x3f
	s.PushLong(int64(uint64(var1) >> bits))
}

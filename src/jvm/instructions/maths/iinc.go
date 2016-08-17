package maths

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * iinc指令给局部变量表中的int变量增加常量
 * 局部变量表索引和常量值都由指令的操作数给出
 */
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	value := localVars.GetInt(self.Index)
	localVars.SetInt(self.Index, value+self.Const)
}

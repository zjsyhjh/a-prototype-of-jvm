package constants

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶
 */
type BIPUSH struct {
	value int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.value = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.value)
	frame.OperandStack().PushInt(i)
}

/*
 * sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶
 */
type SIPUSH struct {
	value int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.value = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.value)
	frame.OperandStack().PushInt(i)
}

package loads

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 加载指令从局部变量表中获取变量，然后推入操作数栈，共33条，按照所操作变量的类型可以分为6类
 */
type ILOAD struct {
	base.Index8Instruction
}

func iload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(value)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	iload(frame, uint(self.Index))
}

/*
 * 加载局部变量表的第1个int型变量进栈
 */
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	iload(frame, 0)
}

/*
 * 加载局部变量表的第2个int型变量进栈
 */
type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	iload(frame, 1)
}

/*
 * 加载局部变量表的第3个int型变量进栈
 */
type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	iload(frame, 2)
}

/*
 * 加载局部变量表的第4个int型变量进栈
 */
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	iload(frame, 3)
}

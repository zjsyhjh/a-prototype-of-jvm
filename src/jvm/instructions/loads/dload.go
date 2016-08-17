package loads

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 加载double型变量入栈
 */
type DLOAD struct {
	base.Index8Instruction
}

func dload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(value)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	dload(frame, uint(self.Index))
}

/*
 * 加载局部变量表的第1个double型变量进栈
 */
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	dload(frame, 0)
}

/*
 * 加载局部变量表的第2个double型变量进栈
 */
type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	dload(frame, 1)
}

/*
 * 加载局部变量表的第3个double型变量进栈
 */
type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	dload(frame, 2)
}

/*
 * 加载局部变量表的第4个double型变量进栈
 */
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	dload(frame, 3)
}

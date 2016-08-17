package loads

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 加载float型变量入栈
 */
type FLOAD struct {
	base.Index8Instruction
}

func fload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(value)
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	fload(frame, uint(self.Index))
}

/*
 * 加载局部变量表的第1个float型变量进栈
 */
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	fload(frame, 0)
}

/*
 * 加载局部变量表的第2个float型变量进栈
 */
type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	fload(frame, 1)
}

/*
 * 加载局部变量表的第3个float型变量进栈
 */
type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	fload(frame, 2)
}

/*
 * 加载局部变量表的第4个float型变量进栈
 */
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	fload(frame, 3)
}

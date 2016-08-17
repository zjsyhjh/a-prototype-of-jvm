package loads

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 加载reference型变量入栈
 */
type ALOAD struct {
	base.Index8Instruction
}

func aload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(value)
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	aload(frame, uint(self.Index))
}

/*
 * 加载局部变量表的第1个reference型变量进栈
 */
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	aload(frame, 0)
}

/*
 * 加载局部变量表的第2个reference型变量进栈
 */
type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	aload(frame, 1)
}

/*
 * 加载局部变量表的第3个reference型变量进栈
 */
type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	aload(frame, 2)
}

/*
 * 加载局部变量表的第4个reference型变量进栈
 */
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	aload(frame, 3)
}

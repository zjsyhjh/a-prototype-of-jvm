package loads

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 加载long型变量入栈
 */
type LLOAD struct {
	base.Index8Instruction
}

func lload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(value)
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	lload(frame, uint(self.Index))
}

/*
 * 加载局部变量表的第1个long型变量进栈
 */
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	lload(frame, 0)
}

/*
 * 加载局部变量表的第2个long型变量进栈
 */
type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	lload(frame, 1)
}

/*
 * 加载局部变量表的第3个long型变量进栈
 */
type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	lload(frame, 2)
}

/*
 * 加载局部变量表的第4个long型变量进栈
 */
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	lload(frame, 3)
}

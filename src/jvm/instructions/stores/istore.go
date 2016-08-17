package stores

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 存储指令和加载指令刚好相反，存储指令是把变量从操作数中弹出，存到局部变量表中
 */
type ISTORE struct {
	base.Index8Instruction
}

func istore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, value)
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	istore(frame, uint(self.Index))
}

/*
 * 存到索引为0的局部变量表中
 */
type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	istore(frame, 0)
}

/*
 * 存到索引为1的局部变量表中
 */
type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	istore(frame, 1)
}

/*
 * 存到索引为2的局部变量表中
 */
type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	istore(frame, 2)
}

/*
 * 存到索引为3的局部变量表中
 */
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	istore(frame, 3)
}

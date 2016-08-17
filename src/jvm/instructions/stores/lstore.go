package stores

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 存储指令和加载指令刚好相反，存储指令是把变量从操作数中弹出，存到局部变量表中
 */
type LSTORE struct {
	base.Index8Instruction
}

func lstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	lstore(frame, uint(self.Index))
}

/*
 * 存到索引为0的局部变量表中
 */
type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	lstore(frame, 0)
}

/*
 * 存到索引为1的局部变量表中
 */
type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	lstore(frame, 1)
}

/*
 * 存到索引为2的局部变量表中
 */
type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	lstore(frame, 2)
}

/*
 * 存到索引为3的局部变量表中
 */
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	lstore(frame, 3)
}

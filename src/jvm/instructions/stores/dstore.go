package stores

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 存储指令和加载指令刚好相反，存储指令是把变量从操作数中弹出，存到局部变量表中
 */
type DSTORE struct {
	base.Index8Instruction
}

func dstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, value)
}

func (self *DSTORE) Execute(frame *rtda.Frame) {
	dstore(frame, uint(self.Index))
}

/*
 * 存到索引为0的局部变量表中
 */
type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	dstore(frame, 0)
}

/*
 * 存到索引为1的局部变量表中
 */
type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	dstore(frame, 1)
}

/*
 * 存到索引为2的局部变量表中
 */
type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	dstore(frame, 2)
}

/*
 * 存到索引为3的局部变量表中
 */
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	dstore(frame, 3)
}

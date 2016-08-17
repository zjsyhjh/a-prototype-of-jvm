package stores

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 存储指令和加载指令刚好相反，存储指令是把变量从操作数中弹出，存到局部变量表中
 */
type ASTORE struct {
	base.Index8Instruction
}

func astore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, value)
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	astore(frame, uint(self.Index))
}

/*
 * 存到索引为0的局部变量表中
 */
type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	astore(frame, 0)
}

/*
 * 存到索引为1的局部变量表中
 */
type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	astore(frame, 1)
}

/*
 * 存到索引为2的局部变量表中
 */
type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	astore(frame, 2)
}

/*
 * 存到索引为3的局部变量表中
 */
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	astore(frame, 3)
}

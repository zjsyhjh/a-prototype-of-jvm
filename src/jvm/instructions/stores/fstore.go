package stores

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 存储指令和加载指令刚好相反，存储指令是把变量从操作数中弹出，存到局部变量表中
 */
type FSTORE struct {
	base.Index8Instruction
}

func fstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, value)
}

func (self *FSTORE) Execute(frame *rtda.Frame) {
	fstore(frame, uint(self.Index))
}

/*
 * 存到索引为0的局部变量表中
 */
type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	fstore(frame, 0)
}

/*
 * 存到索引为1的局部变量表中
 */
type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	fstore(frame, 1)
}

/*
 * 存到索引为2的局部变量表中
 */
type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	fstore(frame, 2)
}

/*
 * 存到索引为3的局部变量表中
 */
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	fstore(frame, 3)
}

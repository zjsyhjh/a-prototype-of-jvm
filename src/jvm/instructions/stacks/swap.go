package stacks

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * swap指令用于交换栈顶两个变量的值
 */
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot1 := s.PopSlot()
	slot2 := s.PopSlot()
	s.PushSlot(slot1)
	s.PushSlot(slot2)
}

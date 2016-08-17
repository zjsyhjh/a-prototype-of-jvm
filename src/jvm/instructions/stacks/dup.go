package stacks

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * dup指令复制栈顶单个数值(数值不能是long或者double类型)，并将复制值压入栈顶
 */
type DUP struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot := s.PopSlot()
	s.PushSlot(slot)
	s.PushSlot(slot)
}

/*
 * dup_x1指令复制栈顶数值(数值不能是long或者double类型)，并将两个复制值压入栈顶
 */
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot1 := s.PopSlot()
	slot2 := s.PopSlot()
	s.PushSlot(slot1)
	s.PushSlot(slot2)
	s.PushSlot(slot1)
}

/*
 * dup_x2指令复制栈顶数值(数值不能是long或者double类型)，并将三个或者两个复制值压入栈顶
 */
type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot1 := s.PopSlot()
	slot2 := s.PopSlot()
	slot3 := s.PopSlot()
	s.PushSlot(slot1)
	s.PushSlot(slot3)
	s.PushSlot(slot2)
	s.PushSlot(slot1)
}

/*
 * dup2指令复制栈顶数值(long或者double类型)，并将复制值压入栈顶
 */
type DUP2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot1 := s.PopSlot()
	slot2 := s.PopSlot()
	s.PushSlot(slot2)
	s.PushSlot(slot1)
	s.PushSlot(slot2)
	s.PushSlot(slot1)
}

/*
 * dup2_x1指令复制栈顶数值(long或者double类型)，并将两个复制值压入栈顶
 */
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot1 := s.PopSlot()
	slot2 := s.PopSlot()
	slot3 := s.PopSlot()
	s.PushSlot(slot2)
	s.PushSlot(slot1)
	s.PushSlot(slot3)
	s.PushSlot(slot2)
	s.PushSlot(slot1)
}

/*
 * dup2_x2指令复制栈顶数值(long或者double类型)，并将三个或者两个复制值压入栈顶
 */
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	slot1 := s.PopSlot()
	slot2 := s.PopSlot()
	slot3 := s.PopSlot()
	slot4 := s.PopSlot()
	s.PushSlot(slot2)
	s.PushSlot(slot1)
	s.PushSlot(slot4)
	s.PushSlot(slot3)
	s.PushSlot(slot2)
	s.PushSlot(slot1)
}

package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * getfield指令获取对象的实例变量值，然后推入操作数栈，需要两个操作数
 * 第一个是uint16索引，第二个是对象引用
 */
type GET_FIELD struct {
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := (frame.Method().Class()).ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()

	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotID := field.SlotID()
	slots := class.StaticVars()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotID))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotID))
	case 'J':
		stack.PushLong(slots.GetLong(slotID))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotID))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotID))
	}
}

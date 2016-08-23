package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * getstatic指令取出类的某个静态变量值，然后推入栈顶
 */
type GET_STATIC struct {
	base.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := (frame.Method().Class()).ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotID := field.SlotID()
	slots := class.StaticVars()
	stack := frame.OperandStack()

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

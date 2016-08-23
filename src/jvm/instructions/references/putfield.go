package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * putfield指令给类的实例变量赋值，需要三个操作数
 * 第一个操作数是uint16索引，来自字节码，通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用
 * 解析这个字段符号引用就可以知道要给类的哪个实例赋值
 * 第二个操作数是变量值
 * 第三个操作数是对象引用，从操作数中弹出
 */
type PUT_FIELD struct {
	base.Index16Instruction
}

/*
 * 拿到当前方法，当前类和当前常量池, 然后解析字段符号引用
 */
func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	/*
	 * 如果是final字段，只能在构造函数中初始化
	 */
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotID := field.SlotID()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		value := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotID, value)
	case 'F':
		value := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")

		}
		ref.Fields().SetFloat(slotID, value)

	case 'J':
		value := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")

		}
		ref.Fields().SetLong(slotID, value)

	case 'D':
		value := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")

		}
		ref.Fields().SetDouble(slotID, value)

	case 'L', '[':
		value := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")

		}
		ref.Fields().SetRef(slotID, value)

	default:

	}
}

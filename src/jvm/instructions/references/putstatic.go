package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * putstatic指令给类的某个静态变量赋值, 需要两个操作数
 * 第一个操作数是uint16索引，来自字节码，通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用
 * 解析这个字段符号引用就可以知道要给类的哪个静态变量赋值
 * 第二个操作数是要赋给静态变量的值，从操作数栈中弹出
 */
type PUT_STATIC struct {
	base.Index16Instruction
}

/*
 * 拿到当前方法，当前类和当前常量池, 然后解析字段符号引用
 */
func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	/*
	 * 如果解析后的字段是实例字段而非静态字段，则抛出异常
	 * 如果是final字段，则实际操作的是静态变量，只能在类初始化中给它赋值(类初始化方法由编译器生成，为<clinit>)
	 */
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	/*
	 * final修饰的字段由类构造器<clinit>方法初始化, 该方法由编译器自动生成
	 */
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotID := field.SlotID()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	/*
	 * 根据字段类型从操作数中弹出相应的值，然后赋值给静态变量
	 */
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotID, stack.PopInt())
	case 'F':
		slots.SetFloat(slotID, stack.PopFloat())
	case 'J':
		slots.SetLong(slotID, stack.PopLong())
	case 'D':
		slots.SetDouble(slotID, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotID, stack.PopRef())
	default:
		//
	}
}

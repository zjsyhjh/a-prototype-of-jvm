package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * 在创建类实例时，编译器会在new指令的后面加入invokespecial指令来调用构造函数初始化对象
 * invokespecial指令用于调用实例构造器<init>方法、私有方法以及父类方法
 * 在编译时选择要调用的方法，它的选择是基于引用类型的(静态绑定), ACC_SUPER除外：在子类调用超类方法的时候发生
 */
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

/*
 * 从方法符号引用中解析出来的类是resolvedClass, 方法是resolvedMethod
 * 如果resolvedMethod是构造函数，则声明resolvedMethod的类必须是resolvedClass
 * 如果resolvedMethod是静态方法，则抛出异常
 */
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	/*
	 * 从操作数栈中得到this引用
	 */
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)

	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	/*
	 * 确保protect方法只能被声明该方法的类或者子类调用
	 */
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	invokedMethod := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() == "<init>" {
		invokedMethod = heap.LookUpMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if invokedMethod == nil || invokedMethod.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, invokedMethod)
}

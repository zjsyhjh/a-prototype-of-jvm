package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * invokeinterface指令调用接口方法，会在运行时再确定一个实现该接口的对象
 * invokeinterface指令的操作码后面跟的是4个字节而非2个字节
 * 前2个字节含义和其他指令相同，是uint16运行时常量池索引
 * 第3个字节是方法传递参数需要的slot数，其含义和Method结构体中的argSlotCount字段相同
 * 第4个字节是留给某些虚拟机使用的，值必须为0
 */
type INVOKE_INTERFACE struct {
	index uint
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()

	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	invokedMethod := heap.LookUpMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if invokedMethod == nil || invokedMethod.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !invokedMethod.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, invokedMethod)
}

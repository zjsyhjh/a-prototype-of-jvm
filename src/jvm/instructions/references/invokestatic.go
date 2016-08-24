package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * invokestatic指令用于调用静态方法
 * 符号解析之后得到方法M，M必须是静态方法，M不能是类初始化方法
 */
type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	class := method.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	base.InvokeMethod(frame, method)
}

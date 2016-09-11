package reserved

/*
 * 没有显示使用lang包用_
 */
import (
	"jvm/instructions/base"
	"jvm/native"
	_ "jvm/native/java/lang"
	_ "jvm/native/sun/misc"
	"jvm/rtda"
)

/*
 * invokenative指令不需要操作数
 */
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)

	if nativeMethod == nil {
		panic("java.lang.UnsatisfiedLinkError : " + className + "~" + methodName + "~" + methodDescriptor)
	}

	nativeMethod(frame)
}

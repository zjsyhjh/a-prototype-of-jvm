package lang

import (
	"jvm/native"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * 注册本地方法
 */
func init() {
	native.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

/*
 * static native Class<?> getPrimitiveClass
 * 先从局部变量表中拿到类名，转成go字符串，最后把类对象引用推入操作数栈
 */
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}

/*
 * private native String getName0
 * 先从局部变量表中拿到this引用，然后与之对应的Class结构体指针
 * 然后拿到类名，转成java字符串并推入操作数栈顶
 */
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

/*
 * private static native boolean
 */
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}

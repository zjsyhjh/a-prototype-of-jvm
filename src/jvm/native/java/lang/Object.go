package lang

import (
	"jvm/native"
	"jvm/rtda"
	"unsafe"
)

/*
 * 注册getClass本地方法
 */
func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

/*
 * public final native Class<?> getClass()
 */
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

/*
 * public native int hashCode()
 * 把对象引用（Object结构体指针）转换成uintptr类型，然后强制转换成int32推入操作数栈顶
 */
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

/*
 * protected native Object clone() throws CloneNotSupportedException;
 * 对象克隆
 */
func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")

	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}

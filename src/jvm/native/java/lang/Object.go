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
 */
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

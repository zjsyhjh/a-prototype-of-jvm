package misc

import (
	"jvm/instructions/base"
	"jvm/native"
	"jvm/rtda"
	"jvm/rtda/heap"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

/*
 * 等价的Java代码
 * public static native void initialize() {
     VM.saveProps.setProperty("foo", "bar");
 }
*/
func initialize(frame *rtda.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	value := heap.JString(vmClass.Loader(), "bar")

	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(value)

	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)
}

/*
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlsClass := classLoader.LoadClass("java/lang/System")
	initClass := jlsClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initClass)
}*/

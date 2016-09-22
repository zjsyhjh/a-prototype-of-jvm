package lang

import (
	"jvm/native"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * 用于记录java虚拟机栈信息
 */
type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

/*
 * private native Throwable fillInStackTrace(int dummy);
 */
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}

	return stes
}

func distanceToObject(class *heap.Class) int {
	var d = 0
	for c := class.SuperClass(); c != nil; c = class.SuperClass() {
		d++
	}
	return d
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()

	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

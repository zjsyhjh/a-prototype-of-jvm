package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
	"reflect"
)

/*
 *ATHROW指令的操作数是一个异常引用对象，从操作数栈弹出
 */
type ATHROW struct {
	base.NoOperandsInstruction
}

func (self *ATHROW) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

/*
 *遍历虚拟机栈，如果当前的栈帧中找不到对应的异常项，则把当前栈帧弹出，继续遍历
 */
func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1
		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)

		if handlerPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()

	jmsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	gomsg := heap.GoString(jmsg)
	println(ex.Class().JavaName() + ": " + gomsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}

package base

import (
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * 方法调用, 在定位到需要调用的方法之后，Java虚拟机要给这个方法创建一个新的栈帧并把它推入虚拟机栈顶
 */
func InvokeMethod(invokeFrame *rtda.Frame, method *heap.Method) {
	thread := invokeFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argsSlotCount := int(method.ArgSlotCount())
	/*
		fmt.Printf("argSlotCount = %v\n", argsSlotCount)
		fmt.Printf(">> pc:%4d %v.%v%v \n", invokeFrame.NextPC(), method.Class().Name(), method.Name(), method.Descriptor())
	*/

	if argsSlotCount > 0 {
		for i := argsSlotCount - 1; i >= 0; i-- {
			slot := invokeFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	/*
		if method.IsNative() {
			if method.Name() == "registerNatives" {
				thread.PopFrame()
			} else {
				panic(fmt.Sprintf("native method: %v.%v%v\n", method.Class().Name(), method.Name(), method.Descriptor()))
			}
		}
	*/
}

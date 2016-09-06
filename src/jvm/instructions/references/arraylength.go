package references

import (
	"fmt"
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * arraylength指令用于获取数组长度
 * 只需一个操作数，从操作数栈中弹出数组引用
 */
type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrayRef := stack.PopRef()
	if arrayRef == nil {
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), frame.Method().Class().Name(), frame.Method().Name(), frame.Method().Descriptor())
		panic("java.lang.NullPointerException")
	}

	arrayLength := arrayRef.ArrayLength()
	stack.PushInt(arrayLength)
}

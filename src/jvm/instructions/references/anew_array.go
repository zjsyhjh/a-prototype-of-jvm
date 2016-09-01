package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * anewarray指令用于创建引用类型数组
 * 该指令需要两个操作数，第一个操作数是一个uint16整型，来自字节码
 * 通过该操作数可以从类的运行时常量池中找到一个类符号引用, 通过解析可以得到数组元素的类
 * 第二个操作数为数组长度，从操作数栈中弹出
 */
type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	classRef := constantPool.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	size := stack.PopInt()
	if size < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrayClass := componentClass.ArrayClass()
	array := arrayClass.NewArray(uint(size))
	stack.PushRef(array)
}

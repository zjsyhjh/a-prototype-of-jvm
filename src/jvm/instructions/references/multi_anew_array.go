package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * mutianewarray指令用于创建多维数组
 * 第一个操作数是个uint16索引， 通过这个索引可以从运行时常量池中找到一个类符号引用
 * 解析这个符号引用可以得到数组类
 * 第二个操作数是一个uint8整数，代表着数组维度
 */
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

/*
 * multianewarray指令需要从操作数栈中弹出n个整数，代表每一维数组长度
 */
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	classRef := constantPool.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrayClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))

	array := newMultiDimensionalArray(arrayClass, counts)
	stack.PushRef(array)
}

/*
 * 从操作数弹出dimensions个数
 */
func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

func newMultiDimensionalArray(arrayClass *heap.Class, counts []int32) *heap.Object {
	count := uint(counts[0])
	array := arrayClass.NewArray(count)

	if len(counts) > 1 {
		refs := array.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(arrayClass.ComponentClass(), counts[1:])
		}
	}
	return array
}

package loads

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * <t>aload系列指令按索引取数组元素值, 然后推入操作数栈中
 */
/*
 * load reference from array
 */
type AALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	refs := arrayRef.Refs()

	checkIndex(len(refs), index)

	stack.PushRef(refs[index])
}

/*
 * load byte from array
 */
type BALOAD struct {
	base.NoOperandsInstruction
}

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	bytes := arrayRef.Bytes()

	checkIndex(len(bytes), index)

	stack.PushInt(int32(bytes[index]))

}

/*
 * load char from array
 */
type CALOAD struct {
	base.NoOperandsInstruction
}

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	chars := arrayRef.Chars()

	checkIndex(len(chars), index)

	stack.PushInt(int32(chars[index]))
}

/*
 * load double from array
 */
type DALOAD struct {
	base.NoOperandsInstruction
}

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	doubles := arrayRef.Doubles()

	checkIndex(len(doubles), index)

	stack.PushDouble(doubles[index])
}

/*
 * load float from array
 */
type FALOAD struct {
	base.NoOperandsInstruction
}

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	floats := arrayRef.Floats()

	checkIndex(len(floats), index)

	stack.PushFloat(floats[index])
}

/*
 * load int from array
 */
type IALOAD struct {
	base.NoOperandsInstruction
}

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	ints := arrayRef.Ints()

	checkIndex(len(ints), index)

	stack.PushInt(ints[index])
}

/*
 * load long from array
 */
type LALOAD struct {
	base.NoOperandsInstruction
}

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	longs := arrayRef.Longs()

	checkIndex(len(longs), index)

	stack.PushLong(longs[index])
}

/*
 * load short from array
 */
type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	shorts := arrayRef.Shorts()

	checkIndex(len(shorts), index)

	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerExecption")
	}
}

func checkIndex(l int, index int32) {
	if index < 0 || index >= int32(l) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

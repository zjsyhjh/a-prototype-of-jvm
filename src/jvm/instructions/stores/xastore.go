package stores

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * <t>astore指令按索引给数组元素复制
 */
/*
 * store value into reference array
 */
type AASTORE struct {
	base.NoOperandsInstruction
}

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	refs := arrayRef.Refs()

	checkIndex(len(refs), index)

	refs[index] = ref
}

/*
 * store value into byte array
 */
type BASTORE struct {
	base.NoOperandsInstruction
}

func (self *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	bytes := arrayRef.Bytes()

	checkIndex(len(bytes), index)

	bytes[index] = int8(value)
}

/*
 * store value into char array
 */
type CASTORE struct {
	base.NoOperandsInstruction
}

func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	chars := arrayRef.Chars()

	checkIndex(len(chars), index)

	chars[index] = uint16(value)
}

/*
 * store value into double array
 */
type DASTORE struct {
	base.NoOperandsInstruction
}

func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	doubles := arrayRef.Doubles()

	checkIndex(len(doubles), index)

	doubles[index] = float64(value)

}

/*
 * store value into float array
 */
type FASTORE struct {
	base.NoOperandsInstruction
}

func (self *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	floats := arrayRef.Floats()

	checkIndex(len(floats), index)

	floats[index] = float32(value)
}

/*
 * store value into int array
 */
type IASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	ints := arrayRef.Ints()

	checkIndex(len(ints), index)

	ints[index] = int32(value)
}

/*
 * store value into long array
 */
type LASTORE struct {
	base.NoOperandsInstruction
}

func (self *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopLong()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	longs := arrayRef.Longs()

	checkIndex(len(longs), index)

	longs[index] = int64(value)
}

/*
 * store value into short array
 */
type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrayRef := stack.PopRef()

	checkNotNil(arrayRef)

	shorts := arrayRef.Shorts()

	checkIndex(len(shorts), index)

	shorts[index] = int16(value)
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

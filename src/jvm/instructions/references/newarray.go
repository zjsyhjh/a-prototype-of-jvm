package references

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * newarray指令用于创建基本数组类型
 * 包括boolean[], byte[], char[], short[], int[], long[], float[], double[]
 * 该指令需要两个操作数，第一个操作数是一个uint8整数，表示要创建的数组类型
 * 第二个操作数是要创建的数组大小， 从操作数栈中弹出
 */
const (
	/*
	 * array type
	 */
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

type NEW_ARRAY struct {
	atype uint8
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

/*
 * 根据atype和size创建数组
 */
func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	size := stack.PopInt()
	if size < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method().Class().Loader()
	arrayClass := getPrimitiveArrayClass(classLoader, self.atype)
	array := arrayClass.NewArray(uint(size))
	stack.PushRef(array)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype")
	}
}

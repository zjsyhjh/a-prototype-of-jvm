package constants

import (
	"jvm/instructions/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

/*
 * ldc从运行时常量池加载常量池
 * ldc和ldc_w指令用于加载int, float和字符串常量，java.lang.Class实例或者MethodType和MethodHandle实例
 * ldc2_w指令用于加载long、double常量
 */
type LDC struct {
	base.Index8Instruction
}

func (self *LDC) Execute(frame *rtda.Frame) {
	ldc(frame, self.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	ldc(frame, self.Index)
}

type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		s.PushLong(c.(int64))
	case float64:
		s.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func ldc(frame *rtda.Frame, index uint) {
	s := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		s.PushInt(c.(int32))
	case float32:
		s.PushFloat(c.(float32))
	/*
	 * 如果ldc从常量池加载字符串常量，则先通过常量拿到Go字符串
	 * 然后把它转成Java字符串实例并推入操作数栈
	 */
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		s.PushRef(internedStr)
	default:
		panic("todo")
	}

}

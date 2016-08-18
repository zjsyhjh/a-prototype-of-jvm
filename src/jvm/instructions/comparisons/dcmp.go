package comparisons

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * double类型比较，结果有4种，大于、等于、小于、无法比较(出现NaN)
 * 当两个double类型变量中至少有一个为NaN时，用fcmpg指令比较结果为1，用fcmpl指令比较结果为-1
 */
type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	dcmp(frame, false)
}

func dcmp(frame *rtda.Frame, flag bool) {
	s := frame.OperandStack()
	var2 := s.PopDouble()
	var1 := s.PopDouble()
	if var1 > var2 {
		s.PushInt(1)
	} else if var1 == var2 {
		s.PushInt(0)
	} else if var1 < var2 {
		s.PushInt(-1)
	} else if flag {
		s.PushInt(1)
	} else {
		s.PushInt(-1)
	}
}

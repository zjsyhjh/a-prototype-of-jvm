package comparisons

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * float类型比较，结果有4种，大于、等于、小于、无法比较(出现NaN)
 * 当两个float类型变量中至少有一个为NaN时，用fcmpg指令比较结果为1，用fcmpl指令比较结果为-1
 */
type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	fcmp(frame, false)
}

func fcmp(frame *rtda.Frame, flag bool) {
	s := frame.OperandStack()
	var2 := s.PopFloat()
	var1 := s.PopFloat()
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

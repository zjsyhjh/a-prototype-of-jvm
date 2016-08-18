package comparisons

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 比较long变量, 将比较结果推入栈顶(1, 0, -1)
 */
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopLong()
	var1 := s.PopLong()
	if var1 > var2 {
		s.PushInt(1)
	} else if var1 == var2 {
		s.PushInt(0)
	} else {
		s.PushInt(-1)
	}
}

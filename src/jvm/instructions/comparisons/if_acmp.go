package comparisons

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * if_acmp<cond>指令
 * if_acmpeq和if_acmpne指令把栈顶两个引用变量弹出，比较是否相同进行跳转
 */
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopRef()
	var1 := s.PopRef()
	if var1 == var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	var2 := s.PopRef()
	var1 := s.PopRef()
	if var1 != var2 {
		base.Branch(frame, self.Offset)
	}
}

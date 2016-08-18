package comparisons

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * if_icmp<cond>指令， 从栈顶弹出2个int值，进行比较，成功则跳转
 */
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if var1, var2 := icmpPopInt(frame); var1 == var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if var1, var2 := icmpPopInt(frame); var1 != var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if var1, var2 := icmpPopInt(frame); var1 < var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if var1, var2 := icmpPopInt(frame); var1 <= var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if var1, var2 := icmpPopInt(frame); var1 > var2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if var1, var2 := icmpPopInt(frame); var1 >= var2 {
		base.Branch(frame, self.Offset)
	}
}

func icmpPopInt(frame *rtda.Frame) (var1, var2 int32) {
	s := frame.OperandStack()
	var2 = s.PopInt()
	var1 = s.PopInt()
	return var1, var2
}

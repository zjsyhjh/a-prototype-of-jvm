package extends

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * ifnull和ifnonnull指令根据引用是否null进行跳转
 */
type IF_NULL struct {
	base.BranchInstruction
}

func (self *IF_NULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

type IF_NONNULL struct {
	base.BranchInstruction
}

func (self *IF_NONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}

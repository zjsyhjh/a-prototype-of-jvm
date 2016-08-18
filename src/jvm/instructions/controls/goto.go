package controls

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * goto指令进行无条件跳转
 */
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

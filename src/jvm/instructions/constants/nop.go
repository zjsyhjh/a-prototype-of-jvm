package constants

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * 常量指令把常量推入操作数栈顶
 * nop指令是最简单的一条指令，它什么都不做
 */
type NOP struct {
	base.NoOperandsInstruction
}

/*
 * do nothing
 */
func (self *NOP) Execute(frame *rtda.Frame) {

}

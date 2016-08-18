package extends

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * goto_w指令和goto指令的区别在于索引从2个字节变成了4个字节
 */
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}

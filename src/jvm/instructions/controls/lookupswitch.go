package controls

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * FetchOperands()方法也要先跳过padding
 * matchOffsets有点像Map，它的key是case值，value是跳转偏移量
 * Execute（）方法先从操作数栈中弹出一个int变量，然后用它查找matchOffsets，看是否能找到匹配的key
 * 如果能, 则按照value给出的值跳转，否则按defaultOffset值跳转
 */
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			base.Branch(frame, int(self.matchOffsets[i+1]))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}

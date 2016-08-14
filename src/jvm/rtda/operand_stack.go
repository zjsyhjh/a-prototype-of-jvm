package rtda

import (
	"math"
)

/*
 * 操作数栈的实现方式和局部变量表类似
 */
type OperandStack struct {
	size  uint
	slots []Slot
}

/*
 * 创建操作数栈
 */
func newOperandStack(maxSize uint) *OperandStack {
	if maxSize > 0 {
		return &OperandStack{
			size:  0,
			slots: make([]Slot, maxSize),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(value int32) {
	self.slots[self.size].num = value
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(value float32) {
	bits := math.Float32bits(value)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

func (self *OperandStack) PushLong(value int64) {
	self.slots[self.size].num = int32(value)
	self.slots[self.size+1].num = int32(value >> 32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (self *OperandStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].reference = ref
	self.size++
}

func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].reference
	self.slots[self.size].reference = nil
	return ref
}

/*
 * 和其他指令不同，栈指令并不关心变量类型
 */
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

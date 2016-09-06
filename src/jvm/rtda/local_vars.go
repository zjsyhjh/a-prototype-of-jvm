package rtda

import (
	"jvm/rtda/heap"
	"math"
)

/*
 * 局部变量表按索引访问，可以看成是一个数组
 * 根据Java虚拟机规范，这个数组中的每个元素至少可以容纳一个int值或者一个引用值
 * 两个连续的元素可以容纳一个long或者double值
 */
type LocalVars []Slot

/*
 * 创建LocalVars实例
 */
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

/*
 * 4种基本类型
 */
func (self LocalVars) SetInt(index uint, value int32) {
	self[index].num = value
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

func (self LocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

/*
 * Long占用2个Slot
 */
func (self LocalVars) SetLong(index uint, value int64) {
	self[index].num = int32(value)
	self[index+1].num = int32(value >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high<<32) | int64(low)
}

/*
 * Double占用2个Slot
 */
func (self LocalVars) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	self[index].num = int32(bits)
	self[index+1].num = int32(bits >> 32)
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

/*
 * 引用类型
 */
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].reference = ref
}

func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].reference
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}

func (self LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}

package heap

func (self *Object) Clone() *Object {
	return &Object{
		class: self.class,
		data:  self.cloneData(),
	}
}

func (self *Object) cloneData() interface{} {
	switch self.data.(type) {
	case []int8:
		element1 := self.data.([]int8)
		element2 := make([]int8, len(element1))
		copy(element2, element1)
		return element2
	case []int16:
		element1 := self.data.([]int16)
		element2 := make([]int16, len(element1))
		copy(element2, element1)
		return element2
	case []uint16:
		element1 := self.data.([]uint16)
		element2 := make([]uint16, len(element1))
		copy(element2, element1)
		return element2
	case []int32:
		element1 := self.data.([]int32)
		element2 := make([]int32, len(element1))
		copy(element2, element1)
		return element2
	case []int64:
		element1 := self.data.([]int64)
		element2 := make([]int64, len(element1))
		copy(element2, element1)
		return element2
	case []float32:
		element1 := self.data.([]float32)
		element2 := make([]float32, len(element1))
		copy(element2, element1)
		return element2
	case []float64:
		element1 := self.data.([]float64)
		element2 := make([]float64, len(element1))
		copy(element2, element1)
		return element1
	case []*Object:
		element1 := self.data.([]*Object)
		element2 := make([]*Object, len(element1))
		copy(element2, element1)
		return element1
	default:
		slot1 := self.data.(Slots)
		slot2 := newSlots(uint(len(slot1)))
		copy(slot2, slot1)
		return slot2
	}
}

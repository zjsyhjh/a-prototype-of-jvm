package heap

/*
 *创建数组对象
 */
func (self *Class) NewArray(size uint) *Object {

	if !self.IsArray() {
		panic("Not array class: " + self.Name())
	}

	switch self.Name() {
	case "[Z":
		return &Object{self, make([]int8, size), nil}
	case "[B":
		return &Object{self, make([]int8, size), nil}
	case "[C":
		return &Object{self, make([]uint16, size), nil}
	case "[S":
		return &Object{self, make([]int16, size), nil}
	case "[I":
		return &Object{self, make([]int32, size), nil}
	case "[J":
		return &Object{self, make([]int64, size), nil}
	case "[F":
		return &Object{self, make([]float32, size), nil}
	case "[D":
		return &Object{self, make([]float64, size), nil}
	default:
		return &Object{self, make([]*Object, size), nil}
	}
}

/*
 * 判断是否是数组
 */
func (self *Class) IsArray() bool {
	return self.className[0] == '['
}

func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.Name())
	return self.classLoader.LoadClass(componentClassName)
}

package heap

/*
 * 方法描述符结构体, 包括参数类型以及返回值类型
 */
type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (self *MethodDescriptor) addParameterType(paramType string) {
	/*
		l := len(self.parameterTypes)
		if l == cap(self.parameterTypes) {
			s := make([]string, l, l+4)
			copy(s, self.parameterTypes)
			self.parameterTypes = s
		}
	*/
	self.parameterTypes = append(self.parameterTypes, paramType)
}

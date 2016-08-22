package heap

func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isAssignableFrom(other *Class) bool {
	if self == other {
		return true
	}
	if !self.IsInterface() {
		return other.IsSubClassOf(self)
	}
	return other.IsImplements(self)
}

func (self *Class) IsImplements(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == other || i.isSubInterfaceOf(other) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(other *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == other || superInterface.isSubInterfaceOf(other) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

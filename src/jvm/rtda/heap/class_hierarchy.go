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

	if !other.IsArray() {
		if !other.IsInterface() {
			// other is Class
			if !self.IsInterface() {
				//self is Class
				return other.IsSubClassOf(self)
			}
			//self is interface
			return other.IsImplements(self)
		} else {
			//other is interface
			// self is not interface
			if !self.IsInterface() {
				return self.isJlObject()
			}
			//self is interface
			return self.isSuperInterfaceOf(other)
		}

	} else {
		//other is array
		if !self.IsArray() {
			if !self.IsInterface() {
				return self.isJlObject()
			}
			return self.isJlCloneable() || self.isJioSerializable()
		} else {
			c1 := other.ComponentClass()
			c2 := self.ComponentClass()
			return c1 == c2 || c2.isAssignableFrom(c1)
		}
	}
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

func (self *Class) isSuperInterfaceOf(other *Class) bool {
	return other.isSubInterfaceOf(self)
}

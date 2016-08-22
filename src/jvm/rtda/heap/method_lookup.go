package heap

/*
 * 从c以及c的超类中查找方法
 */
func LookUpMethodInClass(class *Class, name, descripotr string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descripotr {
				return method
			}
		}
	}
	return nil
}

/*
 * 从c实现的接口中查找
 */
func LookUpMethodInInterface(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := LookUpMethodInInterface(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

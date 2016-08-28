package heap

import (
	"jvm/classfile"
)

/*
 * 接口方法符号引用
 */
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	interfaceMethodRef := &InterfaceMethodRef{}
	interfaceMethodRef.cp = cp
	interfaceMethodRef.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return interfaceMethodRef
}

func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethod()
	}
	return self.method
}

/*
 * class1想通过接口方法符号引用访问class2的某个方法
 * 先解析符号引用得到class2，如果class2不是接口，则抛出异常
 * 否则根据方法名和描述符查找方法, 找不到方法则抛出异常
 * 否则检查class1是否有权限访问该方法
 */
func (self *InterfaceMethodRef) resolveInterfaceMethod() {
	class1 := self.cp.class
	class2 := self.ResolvedClass()

	if !class2.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	interfaceMethod := lookUpInterfaceMethod(class2, self.name, self.descriptor)
	if interfaceMethod == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !interfaceMethod.isAccessibleTo(class1) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = interfaceMethod
}

func lookUpInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return LookUpMethodInInterface(iface.interfaces, name, descriptor)
}

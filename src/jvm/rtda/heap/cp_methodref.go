package heap

import (
	"jvm/classfile"
)

/*
 * 非接口方法符号引用
 */
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	methodRef := &MethodRef{}
	methodRef.cp = cp
	methodRef.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return methodRef
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethod()
	}
	return self.method
}

/*
 * class1想通过方法符号引用访问class2的某个方法
 * 先解析符号引用得到class2，如果class2是接口，则抛出异常
 * 否则根据方法名和描述符查找方法, 找不到方法则抛出异常
 * 否则检查class1是否有权限访问该方法
 */
func (self *MethodRef) resolveMethod() {
	class1 := self.cp.class
	class2 := self.ResolvedClass()

	if class2.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookUpMethod(class2, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(class1) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookUpMethod(class *Class, name, descriptor string) *Method {
	method := LookUpMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookUpMethodInInterface(class.interfaces, name, descriptor)
	}
	return method
}

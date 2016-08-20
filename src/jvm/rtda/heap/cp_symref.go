package heap

/*
 * 符号引用
 * cp保存符号引用所在的运行时常量池指针
 */
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

/*
 * 如果类符号引用已经解析，则直接返回，否则进行解析
 */
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	class1 := self.cp.class
	class2 := class1.classLoader.LoadClass(self.className)
	if !class2.isAccessibleTo(class1) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = class2
}

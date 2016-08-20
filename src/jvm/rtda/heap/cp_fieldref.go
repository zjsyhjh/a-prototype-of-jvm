package heap

import "jvm/classfile"

/*
 * 字段符号引用
 */
type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	fieldRef := &FieldRef{}
	fieldRef.cp = cp
	fieldRef.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return fieldRef
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	class1 := self.cp.class
	class2 := self.ResolvedClass()

	field := lookUpField(class2, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.isAccessibleTo(class1) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

/*
 * 字段查找，从当前类，接口以及超类中查找
 */
func lookUpField(class *Class, name string, desciptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == desciptor {
			return field
		}
	}

	for _, c := range class.interfaces {
		if field := lookUpField(c, name, desciptor); field != nil {
			return field
		}
	}

	if class.superClass != nil {
		return lookUpField(class.superClass, name, desciptor)
	}

	return nil
}

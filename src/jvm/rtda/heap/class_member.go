package heap

import (
	"jvm/classfile"
)

/*
 * 用一个结构体ClassMember来代表字段和方法的信息
 * class字段存放指向Class结构体的指针，通过该字段可以获得字段或者方法所在的类
 * ClassMember结构体定义了基本信息，供Field和Method结构体使用
 */
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

/*
 * 从classfile中复制信息
 */
func (self *ClassMember) copyClassMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

/*
 * 访问标志, 字段和方法共同拥有的
 */
func (self *ClassMember) IsPublic() bool {
	return (self.accessFlags & ACC_PUBLIC) != 0
}

func (self *ClassMember) IsPrivate() bool {
	return (self.accessFlags & ACC_PRIVATE) != 0
}

func (self *ClassMember) IsProtected() bool {
	return (self.accessFlags & ACC_PROTECTED) != 0
}

func (self *ClassMember) IsStatic() bool {
	return (self.accessFlags & ACC_STATIC) != 0
}

func (self *ClassMember) IsFinal() bool {
	return (self.accessFlags & ACC_FINAL) != 0
}

func (self *ClassMember) IsSynthetic() bool {
	return (self.accessFlags & ACC_SYNTHETIC) != 0
}

/*
 * 别的类是否能够操作当前类的成员
 * 如果字段是public，则所有类都可以访问
 * 如果是protected，则只有子类和同一个包下的类能够访问
 * 如果字段有默认访问权限(非public、protected、private), 则只有一个包下的类可以访问
 * 否则字段是private，只有声明这个字段的类可以访问
 */
func (self *ClassMember) isAccessibleTo(other *Class) bool {
	if self.IsPublic() {
		return true
	}
	cur := self.class
	if self.IsProtected() {
		return other == cur || other.isSubClassOf(cur) || cur.getPackageName() == other.getPackageName()
	}

	if !self.IsPrivate() {
		return cur.getPackageName() == other.getPackageName()
	}

	return cur == other
}

/*
 * 返回信息
 */
func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}

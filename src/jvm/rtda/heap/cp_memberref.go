package heap

import (
	"jvm/classfile"
)

/*
 * 与ClassMember结构体类似
 */
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(memberInfo *classfile.ConstantMemberRefInfo) {
	self.className = memberInfo.ClassName()
	self.name, self.descriptor = memberInfo.NameAndTypeDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}

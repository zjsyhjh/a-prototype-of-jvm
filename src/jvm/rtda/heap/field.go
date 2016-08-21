package heap

import (
	"jvm/classfile"
)

/*
 * Field结构体包含了ClassMember的信息
 */
type Field struct {
	ClassMember
	slotID             uint
	constantValueIndex uint
}

func newFields(class *Class, fieldMembers []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(fieldMembers))
	for i, memberInfo := range fieldMembers {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyClassMemberInfo(memberInfo)
		fields[i].copyAttributes(memberInfo)
	}
	return fields
}

func (self *Field) copyAttributes(fieldMember *classfile.MemberInfo) {
	if attr := fieldMember.ConstantValueAttribute(); attr != nil {
		self.constantValueIndex = uint(attr.ConstantValueIndex())
	}
}

/*
 * 返回信息
 */
func (self *Field) SlotID() uint {
	return self.slotID
}

func (self *Field) ConstantValueIndex() uint {
	return self.constantValueIndex
}

/*
 * 字段特有的访问标志
 */
func (self *Field) IsVolatile() bool {
	return (self.accessFlags & ACC_VOLATILE) != 0
}

func (self *Field) IsTransient() bool {
	return (self.accessFlags & ACC_TRANSIENT) != 0
}

func (self *Field) IsEnum() bool {
	return (self.accessFlags & ACC_ENUM) != 0
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

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

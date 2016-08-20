package heap

import (
	"jvm/classfile"
)

/*
 * 方法符号引用
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

func (self *MethodRef) resolveMethod() {
	//todo
}

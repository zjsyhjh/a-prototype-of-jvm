package heap

import (
	"jvm/classfile"
)

/*
 * 类符号引用
 */
type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	classRef := &ClassRef{}
	classRef.cp = cp
	classRef.className = classInfo.Name()
	return classRef
}

package classfile

/*
 * CONSTANT_Fieldref_info表示字段符号引用, CONSTANT_Methodref_info表示普通方法引用符号,
 * CONSTANT_InterfaceMethodref_info表示接口方法引用符号
 CONSTANT_Fieldref_info {
     u1 tag;
     u2 class_index;
     u2 name_and_type_index;
 }

 CONSTANT_Methodref_info {
     u1 tag;
     u2 class_index;
     u2 name_and_type_index;
 }

 CONSTANT_InterfaceMethodref_info {
     u1 tag;
     u2 class_index;
     u2 name_and_type_index;
 }
*/
type ConstantMemberRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
	constantPool     ConstantPool
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (cmri *ConstantMemberRefInfo) readInfo(cr *ClassReader) {
	cmri.classIndex = cr.readUint16()
	cmri.nameAndTypeIndex = cr.readUint16()
}

func (cmri *ConstantMemberRefInfo) ClassName() string {
	return cmri.constantPool.getUtf8(cmri.classIndex)
}

func (cmri *ConstantMemberRefInfo) NameAndTypeDescriptor() string {
	return cmri.constantPool.getUtf8(cmri.nameAndTypeIndex)
}

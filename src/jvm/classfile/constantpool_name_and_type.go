package classfile

/*
 * CONSTANT_NameAndType_info给出字段或方法的名称和描述符
 CONSTANT_NameAndType_info {
     u1 tag;
     u2 name_index;
     u2 descriptor_index;
 }
 * 字段和方法名就是代码中出现的（或者编译器生成的）字段和方法的名字
 * 字段或方法名由name_index给出，字段或方法的描述符由descriptor_index给出
 * name_index, descriptor_index都是常量池索引，指向CONSTANT_Utf8_info常量
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cnati *ConstantNameAndTypeInfo) readInfo(cr *ClassReader) {
	cnati.nameIndex = cr.readUint16()
	cnati.descriptorIndex = cr.readUint16()
}

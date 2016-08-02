package classfile

/*
 * CONSTANT_Class_info常量表示类或者接口的符号引用
 CONSTANT_Class_info {
     u1 tag;
     u2 name_index;
 }
 * 和CONSTANT_String_info类似，name_index是常量池索引，指向CONSTANT_Utf8_info常量
*/
type ConstantClassInfo struct {
	nameIndex    uint16
	constantPool ConstantPool
}

func (cci *ConstantClassInfo) readInfo(cr *ClassReader) {
	cci.nameIndex = cr.readUint16()
}

func (cci *ConstantClassInfo) Name() string {
	return cci.constantPool.getUtf8(cci.nameIndex)
}

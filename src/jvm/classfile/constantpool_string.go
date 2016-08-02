package classfile

/*
 * CONSTANT_String_info表示java.lang.String字符串常量
 CONSTANT_String_info {
     u1 tag;
     u2 string_index;
 }
 * CONSTANT_String_info本身并不存放字符串数据，只存放常量池索引，这个索引指向一个CONSTANT_Utf8_info
*/
type ConstantStringInfo struct {
	stringIndex  uint16
	constantPool ConstantPool
}

func (csi *ConstantStringInfo) readInfo(cr *ClassReader) {
	csi.stringIndex = cr.readUint16()
}

func (csi *ConstantStringInfo) String() string {
	return csi.constantPool.getUtf8(csi.stringIndex)
}

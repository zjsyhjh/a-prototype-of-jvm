package classfile

/*
 * SourceFile是可选定长属性，只会出现在ClassFile结构中，用于指出源文件名
 SourceFile_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
     u2 sourcefile_index;
 }
 * attribute_length的值必须为2
*/
type SourceFileAttribute struct {
	constantPool    ConstantPool
	sourceFileIndex uint16
}

func (sfa *SourceFileAttribute) readInfo(cr *ClassReader) {
	sfa.sourceFileIndex = cr.readUint16()
}

func (sfa *SourceFileAttribute) SourceFileName() string {
	return sfa.constantPool.getUtf8(sfa.sourceFileIndex)
}

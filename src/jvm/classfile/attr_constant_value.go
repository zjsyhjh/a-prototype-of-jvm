package classfile

/*
 * ConstantValue是定长属性，只会出现在field_info结构中，用于表示常量表达式的值
 ConstantValue_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
     u2 constantvalue_index;
 }
 * attribute_length的值必须为2
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (cva *ConstantValueAttribute) readInfo(cr *ClassReader) {
	cva.constantValueIndex = cr.readUint16()
}

func (cva *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return cva.constantValueIndex
}

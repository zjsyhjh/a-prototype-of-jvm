package classfile

/*
 * 目前一共14种属性，定义一个AttributeInfo接口，需要具体的属性实现
 * 属性是可以扩展的，而常量是严格定义的，这也是不使用tag而用属性名来区别不同的属性的原因
 * 不同的虚拟机可以实现自己定义的属性类型
 attribute_info {
     u2 attribute_name_index;
     u4 attribute_length;
     u1 info[attribute_length];
 }
*/
type AttributeInfo interface {
	readInfo(cr *ClassReader)
}

func readAttributes(cr *ClassReader, cp ConstantPool) []AttributeInfo {
	count := cr.readUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = readAttribute(cr, cp)
	}
	return attributes
}

func readAttribute(cr *ClassReader, cp ConstantPool) AttributeInfo {
	attributeNameIndex := cr.readUint16()
	attributeName := cp.getUtf8(attributeNameIndex)
	attributeLength := cr.readUint32()
	attributeInfo := newAttributeInfo(attributeName, attributeLength, cp)
	attributeInfo.readInfo(cr)
	return attributeInfo
}

/*
 * 这里暂时只实现8种属性
 */
func newAttributeInfo(attributeName string, attributeLength uint32, cp ConstantPool) AttributeInfo {
	switch attributeName {
	case "Code":
		return &CodeAttribute{constantPool: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{constantPool: cp}
	default:
		return &UnparsedAttribute{attributeName, attributeLength, nil}
	}
}

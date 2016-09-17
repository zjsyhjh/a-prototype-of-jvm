package classfile

/*
 * Code是变长属性，只存在于method_info结构中,Code属性中存放字节码等方法相关信息
 * max_stack给出操作数栈的最大深度，max_locals给出局部变量表的大小
 Code_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
     u2 max_stack;
     u2 max_locals;
     u4 code_length;
     u1 code[code_length];
     {
         u2 start_pc;
         u2 end_pc;
         u2 handle_pc;
         u2 catch_type;
     } exception_table[exception_table_length];
     u2 attributes_count;
     attribute_info attributes[attributes_count];
 }
*/
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlePc  uint16
	catchType uint16
}

func readExceptionTable(cr *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := cr.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   cr.readUint16(),
			endPc:     cr.readUint16(),
			handlePc:  cr.readUint16(),
			catchType: cr.readUint16(),
		}
	}
	return exceptionTable
}

func (ete *ExceptionTableEntry) StartPc() uint16 {
	return ete.startPc
}

func (ete *ExceptionTableEntry) EndPc() uint16 {
	return ete.endPc
}

func (ete *ExceptionTableEntry) HandlePc() uint16 {
	return ete.handlePc
}

func (ete *ExceptionTableEntry) CatchType() uint16 {
	return ete.catchType
}

/*
 * CodeAttribute
 */
type CodeAttribute struct {
	constantPool   ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (ca *CodeAttribute) readInfo(cr *ClassReader) {
	ca.maxStack = cr.readUint16()
	ca.maxLocals = cr.readUint16()
	codeLength := cr.readUint32()
	ca.code = cr.readBytes(codeLength)
	ca.exceptionTable = readExceptionTable(cr)
	ca.attributes = readAttributes(cr, ca.constantPool)
}

func (ca *CodeAttribute) MaxStack() uint {
	return uint(ca.maxStack)
}

func (ca *CodeAttribute) MaxLocals() uint {
	return uint(ca.maxLocals)
}

func (ca *CodeAttribute) Code() []byte {
	return ca.code
}

func (ca *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range ca.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

func (ca *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return ca.exceptionTable
}

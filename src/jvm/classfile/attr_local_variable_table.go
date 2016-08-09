package classfile

/*
 * LocalVariableTable属性表中存放方法的局部变量信息
 LocalVariableTable_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
     u2 local_variable_table_length;
     {
         u2 start_pc;
         u2 length;
         u2 name_index;
         u2 descriptor_index;
         u2 index;
     } local_variable_table[local_variable_table_length];
 }
*/
type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

func (lvta *LocalVariableTableAttribute) readInfo(cr *ClassReader) {
	localVariableTableLength := cr.readUint16()
	lvta.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range lvta.localVariableTable {
		lvta.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         cr.readUint16(),
			length:          cr.readUint16(),
			nameIndex:       cr.readUint16(),
			descriptorIndex: cr.readUint16(),
			index:           cr.readUint16(),
		}
	}
}

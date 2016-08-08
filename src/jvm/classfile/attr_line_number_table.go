package classfile

/*
 * LineNumberTable属性表用于存放方法的行信息，属于调试信息，不是运行时必需
 LineNumberTable_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
     u2 line_number_table_length;
     {
         u2 start_pc;
         u2 line_number;
     } line_number_table[line_number_table_length];
 }
*/
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (lnta *LineNumberTableAttribute) readInfo(cr *ClassReader) {
	lineNumberTableLength := cr.readUint16()
	lnta.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lnta.lineNumberTable {
		lnta.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    cr.readUint16(),
			lineNumber: cr.readUint16(),
		}
	}
}

func (lnta *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lnta.lineNumberTable) - 1; i >= 0; i-- {
		lineNumberTableEntry := lnta.lineNumberTable[i]
		if pc >= int(lineNumberTableEntry.startPc) {
			return int(lineNumberTableEntry.lineNumber)
		}
	}
	return -1
}

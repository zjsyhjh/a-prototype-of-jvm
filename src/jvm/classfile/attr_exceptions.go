package classfile

/*
 * Exceptions_attribute是变长属性，记录方法抛出的异常表
 Exceptions_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
     u2 number_of_exceptions;
     u2 exception_index_table[number_of_exceptions];
 }
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (ea *ExceptionsAttribute) readInfo(cr *ClassReader) {
	ea.exceptionIndexTable = cr.readUint16Table()
}

func (ea *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return ea.exceptionIndexTable
}

package classfile

/*
 * attribute_info {
     u2 attribute_name_index;
     u4 attribute_length;
     u1 info[attribute_length];
 }
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (ua *UnparsedAttribute) readInfo(cr *ClassReader) {
	ua.info = cr.readBytes(ua.length)
}

func (ua *UnparsedAttribute) Info() []byte {
	return ua.info
}

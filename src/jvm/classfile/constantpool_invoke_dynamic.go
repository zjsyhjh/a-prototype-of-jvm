package classfile

/*
 *
 CONSTANT_MethodHandle_info {
     u1 tag;
     u2 reference_kind;
     u2 reference_index;
 }
*/
type ConstantMethodHandleInfo struct {
	referencekind  uint8
	referenceIndex uint16
}

func (cmhi *ConstantMethodHandleInfo) readInfo(cr *ClassReader) {
	cmhi.referencekind = cr.readUint8()
	cmhi.referenceIndex = cr.readUint16()
}

/*
 *
 CONSTANT_MethodType_info {
     u1 tag;
     u2 descriptor_index;
 }
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (cmti *ConstantMethodTypeInfo) readInfo(cr *ClassReader) {
	cmti.descriptorIndex = cr.readUint16()
}

/*
 * CONSTANT_InvokeDynamic_info {
     u1 tag;
     u2 bootstrap_method_attr_index;
     u2 name_and_type_index;
 }
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (cidi *ConstantInvokeDynamicInfo) readInfo(cr *ClassReader) {
	cidi.bootstrapMethodAttrIndex = cr.readUint16()
	cidi.nameAndTypeIndex = cr.readUint16()
}

package classfile

/*
 * class文件中字段和方法的结构类似
 field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
 }
 method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
 }
 *
*/
/*
 * 给出统一定义
 */
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

/*
 * 取得字段或者方法信息表
 */
func readMembers(cr *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := cr.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members {
		members[i] = readMember(cr, cp)
	}

	return members
}

/*
 * 取得某个字段或者方法信息
 */
func readMember(cr *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

/*
 * 取得访问标志符
 */
func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

/*
 * 取得字段或者方法名
 */
func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

/*
 * 取得字段或者方法描述符
 */
func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

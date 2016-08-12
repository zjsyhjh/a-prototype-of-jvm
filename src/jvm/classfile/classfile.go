package classfile

import (
	"fmt"
)

/*
 * 在Java虚拟机中，整个class文件被描述为一个classfile结构，如下所示
 * u2, u4分别代表2个字节的无符号整型和4个字节的无符号整型，对应golang中的uint16和uint32
 ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
 *
*/

/*
 * ClassFile结构体如实反映了Java虚拟机定义的class文件格式
 * 格式为：魔数、次版本号、主版本号、常量池、类访问标志、类名、超类名的索引、接口索引表、字段、方法名、属性表
 */
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/*
 * 解析Class文件
 */
func ParseClassFile(classData []byte) (cf *ClassFile, err error) {
	/*
	 * golang中抛出的一个panic异常，可以在defer中通过recover捕获，然后正常处理
	 */
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}

	cf.readClassFile(cr)

	return cf, nil
}

func (cf *ClassFile) readClassFile(cr *ClassReader) {
	cf.readAndCheckMagic(cr)
	cf.readAndCheckVersion(cr)
	cf.constantPool = getConstantPool(cr)
	cf.accessFlags = cr.readUint16()
	cf.thisClass = cr.readUint16()
	cf.superClass = cr.readUint16()
	cf.interfaces = cr.readUint16Table()
	cf.fields = readMembers(cr, cf.constantPool)
	cf.methods = readMembers(cr, cf.constantPool)
	cf.attributes = readAttributes(cr, cf.constantPool)
}

/*
 * Class文件的开始4个字节为魔数，"0xCAFEBABE"
 */
func (cf *ClassFile) readAndCheckMagic(cr *ClassReader) {
	cf.magic = cr.readUint32()
	if cf.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
 * 获取次版本号以及主版本号
 * jdk1.0.2 class文件版本号为45.0~45.3
 * jdk1.1 class文件版本号为45.0~45.65535
 * jdk1.2 ~ jdk8分别对应46.0, 47.0, 48.0, 49.0, 50.0, 51.0, 52.0
 * 特定的Java虚拟机只支持版本号范围内的class文件， 参考Java SE 8，支持版本号范围为45.0~52.0
 */
func (cf *ClassFile) readAndCheckVersion(cr *ClassReader) {
	cf.minorVersion = cr.readUint16()
	cf.majorVersion = cr.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

/*
 * 返回次版本号
 */
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

/*
 * 返回主版本号
 */
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

/*
 * 返回常量池
 */
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

/*
 * 返回类访问标志
 */
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

/*
 * 返回字段
 */
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

/*
 * 返回方法
 */
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

/*
 * 返回类名
 */
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

/*
 * 返回父类名, java.lang.Object没有父类，superclass为0
 */
func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

/*
 * 返回接口名
 */
func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, interfaceNameIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(interfaceNameIndex)
	}
	return interfaceNames
}

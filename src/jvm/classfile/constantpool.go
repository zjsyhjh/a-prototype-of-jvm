package classfile

import (
	"strconv"
)

/*
 * 常量池也是一个表，如果表头给出的值为n，则有效索引为1~n-1, 0无效
 * CONSTANT_Long_info和CONSTANT_Double_info各占2个位置
 */
type ConstantPool struct {
	constantPool []ConstantInfo
}

/*
 * 取得常量池
 */
func getConstantPool(cr *ClassReader) ConstantPool {
	constantPool := ConstantPool{}

	size := cr.readUint16()
	constantPool.constantPool = make([]ConstantInfo, size)

	for i := 1; i < size; i++ {
		constantPool.constantPool[i] = readConstantInfo(cr, constantPool)
		switch constantPool.constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return constantPool.constantPool
}

/*
 * 取得下标为index的常量信息
 */
func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if constantInfo := cp.constantPool[index]; constantInfo != nil {
		return constantInfo
	}
	panic("Invalid constant pool index: " + strconv.Itoa(index))
}

/*
 * 取得下标为index的常量的字段或方法的名称和描述符
 */
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	constantInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	_name := cp.getUtf8(constantInfo.nameIndex)
	_type := cp.getUtf8(constantInfo.descriptorIndex)
	return _name, _type
}

/*
 * 取得下标为index的类名
 */
func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

/*
 * 取得下标为index的常量池中的字符串
 */
func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.value
}

package heap

import (
	"fmt"
	"jvm/classfile"
)

/*
 * 运行时常量池主要存放两类信息: 字面量和符号引用
 * 字面量包括整数、浮点数和字符串字面量
 * 符号引用包括类符号引用、字段符号引用、方法符号引用和接口方法符号引用
 */
type Constant interface{}

type ConstantPool struct {
	class     *Class
	constants []Constant
}

/*
 * newConstantPool方法把class文件中的常量池转化成运行时常量池
 */
func newConstantPool(class *Class, cfconstantPool classfile.ConstantPool) *ConstantPool {
	count := len(cfconstantPool.ConstantPool())
	constants := make([]Constant, count)
	cp := &ConstantPool{class, constants}

	for i := 1; i < count; i++ {
		cpInfo := cfconstantPool.ConstantPool()[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			constants[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			constants[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			constants[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			constants[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			constants[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			constants[i] = newClassRef(cp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldRefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			constants[i] = newFieldRef(cp, fieldRefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			constants[i] = newMethodRef(cp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			interfaceMethodRefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			constants[i] = newInterfaceMethodRef(cp, interfaceMethodRefInfo)
		default:
			//todo
		}
	}
	return cp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if constant := self.constants[index]; constant != nil {
		return constant
	}
	panic(fmt.Sprintf("No constant at the index %d\n", index))
}

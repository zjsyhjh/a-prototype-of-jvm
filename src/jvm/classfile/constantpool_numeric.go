package classfile

import (
	"math"
)

/*
 * 该文件中定义的结构体都实现了ConstantInfo接口
 */

/*
 * CONSTANT_Integer_info使用4个字节存储整数常量
 CONSTANT_Integer_info {
     u1 tag;
     u4 bytes;
 }
*/
type ConstantIntegerInfo struct {
	value int32
}

func (cii *ConstantIntegerInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	cii.value = int32(bytes)
}

func (cii *ConstantIntegerInfo) Value() int32 {
	return cii.value
}

/*
 * CONSTANT_Float_info使用4个字节存储IEEE754单精度浮点数常量
 CONSTANT_Float_info {
     u1 tag;
     u4 bytes;
 }
*/
type ConstantFloatInfo struct {
	value float32
}

func (cfi *ConstantFloatInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	cfi.value = math.Float32frombits(bytes)
}

func (cfi *ConstantFloatInfo) Value() float32 {
	return cfi.value
}

/*
 * CONSTANT_Long_info使用8个字节存储整型常量
 CONSTANT_Long_info {
     u1 tag;
     u4 high_bytes;
     u4 low_bytes;
 }
*/
type ConstantLongInfo struct {
	value int64
}

func (cli *ConstantLongInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	cli.value = int64(bytes)
}

func (cli *ConstantLongInfo) Value() int64 {
	return cli.value
}

/*
 * CONSTANT_Double_info使用8个字节存储IEEE754双精度浮点数
 CONSTANT_Double_info {
     u1 tag;
     u4 high_bytes;
     u4 low_bytes;
 }
*/
type ConstantDoubleInfo struct {
	value float64
}

func (cdi *ConstantDoubleInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	cdi.value = math.Float64frombits(bytes)
}

func (cdi *ConstantDoubleInfo) Value() float64 {
	return cdi.value
}

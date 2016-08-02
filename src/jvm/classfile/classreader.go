package classfile

import (
	"encoding/binary"
)

/*
 * 使用一个结构体来读取class文件内容
 */
type ClassReader struct {
	data []byte
}

/*
 * 读取一个u1类型的数据
 */
func (cr *ClassReader) readUint8() uint8 {
	value := cr.data[0]
	cr.data = cr.data[1:]
	return value
}

/*
 * 读取一个u2类型的数据, JVM采用大端字节序，即最高位自己存放在最低的内存地址处
 */
func (cr *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return value
}

/*
 * 读取一个u4类型的数据
 */
func (cr *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return value
}

/*
 *读取一个u8类型的数据，不过Java虚拟机并没有定义u8类型的数据
 */
func (cr *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return value
}

/*
 * 读取一个uint16表，表的大小由开头的uint16数据指出
 */
func (cr *ClassReader) readUint16Table() []uint16 {
	size := cr.readUint16()
	table := make([]uint16, size)
	for i := range table {
		table[i] = cr.readUint16()
	}
	return table
}

/*
 * 读取指定数目的字节
 */
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}

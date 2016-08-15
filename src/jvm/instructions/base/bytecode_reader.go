package base

/*
 * BytecodeReader用于读取字节码
 * code字段存放字节码，pc字段记录读取到了哪个字节
 */
type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

/*
 * call by tableswitch and lookupswitch
 */
func (self *BytecodeReader) ReadInt32s(size int32) []int32 {
	ints := make([]int32, size)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

/*
 * call by tableswitch and lookupswitch
 * tableswitch指令操作码的后面有0~3个字节的padding， 以保证defaultOffset在字节码中的地址是4的倍数
 */
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}

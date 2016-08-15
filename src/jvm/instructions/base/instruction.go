package base

import (
	"jvm/rtda"
)

/*
 * base目录下的代码所有的指令共用
 * Java虚拟机解释器的大致逻辑如下: 计算PC、指令解码、指令执行
 for {
     pc := calculatePC()
     opcode := bytecode[pc]
     inst := createInst(opcode)
     inst.fetchOperands(opcode)
     inst.execute()
 }
 * fetchOperands()方法从字节码中提取操作数，execute()方法执行指令逻辑
*/
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

/*
 * NoOperandsInstruction代表没有操作数的指令, 因此没有定义任何字段
 */
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

/*
 * BranchInstruction表示跳转指令，Offset字段存放跳转偏移量
 */
type BranchInstruction struct {
	Offset int
}

/*
 * 从字节码中读取一个int16的整数，转成int之后赋值给Offset
 */
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*
 * 存储和加载指令都需要根据索引存取局部变量表，索引由单字节操作数给出
 */
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/*
 * 一些指令需要访问运行时常量池，常量池索引由两字节操作数给出
 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

package constants

/*
 * const系列指令把隐藏在操作码中的常量值推入操作数栈顶
 */

import (
	"jvm/instructions/base"
	"jvm/rtda"
)

/*
 * aconst_null指令把null引用推入操作数栈顶
 */
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

/*
 * dconst_0指令把double型0推入操作数栈顶
 */
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

/*
 * dconst_1指令把double型1推入操作数栈顶
 */
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

/*
 * fconst_0指令把float型0推入操作数栈顶
 */
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

/*
 * fconst_1指令把float型1推入操作数栈顶
 */
type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

/*
 * fconst_2指令把float型2推入操作数栈顶
 */
type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

/*
 * iconst_m1指令把int型-1推入操作数栈顶
 */
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

/*
 * iconst_0指令把int型0推入操作数栈顶
 */
type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

/*
 * iconst_1指令把int型1推入操作数栈顶
 */
type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

/*
 * iconst_2指令把int型0推入操作数栈顶
 */
type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

/*
 * iconst_3指令把int型3推入操作数栈顶
 */
type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

/*
 * iconst_4指令把int型4推入操作数栈顶
 */
type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

/*
 * iconst_5指令把int型5推入操作数栈顶
 */
type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

/*
 * lconst_0指令把long型0推入操作数栈顶
 */
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

/*
 * lconst_1指令把long型1推入操作数栈顶
 */
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

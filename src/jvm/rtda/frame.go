package rtda

import (
	"jvm/rtda/heap"
)

/*
 * 栈帧中保存方法执行的状态，包括局部变量表(Local Variable)和操作数栈(Operand Stack)等
 * 用链表来实现栈
 */
type Frame struct {
	localVars    LocalVars
	operandStack *OperandStack
	lower        *Frame
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

/*
 * 取得局部变量表
 */
func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}

/*
 * 取得操作数栈
 */
func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

/*
 * 取得线程
 */
func (frame *Frame) Thread() *Thread {
	return frame.thread
}

/*
 * 取得方法
 */
func (frame *Frame) Method() *heap.Method {
	return frame.method
}

/*
 * 返回下一个PC值
 */
func (frame *Frame) NextPC() int {
	return frame.nextPC
}

/*
 * 设置PC值
 */
func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}

func (frame *Frame) RevertNextPC() {
	frame.nextPC = frame.thread.pc
}

package rtda

/*
 * 栈帧中保存方法执行的状态，包括局部变量表(Local Variable)和操作数栈(Operand Stack)等
 * 用链表来实现栈
 */
type Frame struct {
	localVars    LocalVars
	operandStack *OperandStack
	lower        *Frame
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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

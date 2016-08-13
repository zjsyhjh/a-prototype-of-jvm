package rtda

/*
 * 栈帧中保存方法执行的状态，包括局部变量表(Local Variable)和操作数栈(Operand Stack)等
 * 用链表来实现栈
 */
type Frame struct {
	localVars    LocalVars
	operandStack *OperandStack
	lower        *Frame
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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

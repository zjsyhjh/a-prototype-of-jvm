package rtda

import "jvm/rtda/heap"

/*
 * 线程私有的运行时数据区用于辅助执行Java字节码
 * 每个线程都有自己的PC寄存器(Program Counter)以及Java虚拟机栈(JVM Stack)
 * PC中存放的是当前正在执行的Java虚拟机指令的地址
 * Java虚拟机栈又由栈帧(Stack Frame)构成
 * 栈帧中保存方法执行的状态，包括局部变量表(Local Variable)和操作数栈(Operand Stack)等
 */
type Thread struct {
	pc    int
	stack *Stack
}

/*
 * 创建Thread实例
 * 其中栈大小可以通过通过命令行参数设定
 */
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

/*
 * 创建栈帧
 */
func (td *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(td, method)
}

/*
 * 返回当前线程PC寄存器的值
 */
func (td *Thread) PC() int {
	return td.pc
}

/*
 * 设置当前线程PC寄存器的值
 */
func (td *Thread) SetPC(pc int) {
	td.pc = pc
}

/*
 * 放入栈帧
 */
func (td *Thread) PushFrame(frame *Frame) {
	td.stack.push(frame)
}

/*
 * 弹出栈帧
 */
func (td *Thread) PopFrame() *Frame {
	return td.stack.pop()
}

/*
 * 获取当前栈帧
 */
func (td *Thread) CurrentFrame() *Frame {
	return td.stack.top()
}

/*
 * 和CurrentFrame一样
 */
func (td *Thread) TopFrame() *Frame {
	return td.stack.top()
}

func (td *Thread) IsStackEmpty() bool {
	return td.stack.isEmpty()
}

func (td *Thread) ClearStack() {
	td.stack.clear()
}

func (td *Thread) GetFrames() []*Frame {
	return td.stack.getFrames()
}

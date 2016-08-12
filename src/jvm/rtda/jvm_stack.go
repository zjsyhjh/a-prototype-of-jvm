package rtda

/*
 * Java虚拟机栈可以是连续的空间，也可以不连续；可以固定大小，也可以在运行时动态扩展
 */
type Stack struct {
	maxSize  uint
	size     uint
	topFrame *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowErrro")
	}

	if s.topFrame != nil {
		frame.lower = s.topFrame
	}

	s.topFrame = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s.topFrame == nil {
		panic("JVM Stack is empty!")
	}

	frame := s.topFrame
	s.topFrame = frame.lower
	frame.lower = nil
	s.size--
	return frame
}

func (s *Stack) top() *Frame {
	if s.topFrame == nil {
		panic("JVM Stack is empty!")
	}
	return s.topFrame
}

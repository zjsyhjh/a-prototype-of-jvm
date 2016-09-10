package heap

import (
	"jvm/classfile"
)

/*
 * maxStack和maxLocals分别存放操作数栈和局部变量表大小，由javac编译器决定
 * code存放字节码
 */
type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, methodMembers []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(methodMembers))
	for i, memberInfo := range methodMembers {
		methods[i] = newMethod(class, memberInfo)
	}
	return methods
}

/*
 * 本地方法在class文件中没有Code属性，如果是本地方法，则注入字节码及其他信息
 */
func newMethod(class *Class, methodMember *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyClassMemberInfo(methodMember)
	method.copyAttributes(methodMember)
	md := parseMethodDescriptor(method.descriptor)
	method.calArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

/*
 * 注入字节码及其他信息
 */
func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1}
	case 'D':
		self.code = []byte{0xfe, 0xaf}
	case 'F':
		self.code = []byte{0xfe, 0xae}
	case 'J':
		self.code = []byte{0xfe, 0xad}
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0}
	default:
		self.code = []byte{0xfe, 0xac}
	}
}

func (self *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	if codeAttr := memberInfo.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

/*
 * 获取方法参数个数，通过解析方法的描述符可以得到
 */
func (self *Method) calArgSlotCount(parameterTypes []string) {
	for _, paramType := range parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}

	/*
	 * this parameter
	 */
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

/*
 * 方法特有的访问标志
 */
func (self *Method) IsSynchronized() bool {
	return (self.accessFlags & ACC_SYNCHRONIZED) != 0
}

func (self *Method) IsBridge() bool {
	return (self.accessFlags & ACC_BRIDGE) != 0
}

func (self *Method) IsVarargs() bool {
	return (self.accessFlags & ACC_VARARGS) != 0
}

func (self *Method) IsNative() bool {
	return (self.accessFlags & ACC_NATIVE) != 0
}

func (self *Method) IsAbstract() bool {
	return (self.accessFlags & ACC_ABSTRACT) != 0
}

func (self *Method) IsStrict() bool {
	return (self.accessFlags & ACC_STRICT) != 0
}

/*
 * 返回信息
 */
func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

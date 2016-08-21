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
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyClassMemberInfo(memberInfo)
		methods[i].copyAttributes(memberInfo)
		methods[i].calArgSlotCount()
	}
	return methods
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
func (self *Method) calArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
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

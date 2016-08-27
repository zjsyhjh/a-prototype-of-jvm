package heap

import "strings"

/*
 * 解析方法描述符
 */
type MethodDescrptorParser struct {
	methodDescriptor string
	offset           int
	parsed           *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescrptorParser{}
	return parser.parse(descriptor)
}

func (self *MethodDescrptorParser) parse(descriptor string) *MethodDescriptor {
	self.methodDescriptor = descriptor
	self.offset = 0
	self.parsed = &MethodDescriptor{}

	self.checkStartParam()
	self.parseParamTypes()
	self.checkEndParam()
	self.parseReturnType()
	self.checkFinish()
	return self.parsed
}

func (self *MethodDescrptorParser) parseParamTypes() {
	for {
		paramType := self.parseFieldType()
		if paramType != "" {
			self.parsed.addParameterType(paramType)
		} else {
			break
		}
	}
}

func (self *MethodDescrptorParser) parseReturnType() {
	if self.readUint8() == 'V' {
		self.parsed.returnType = "V"
		return
	}
	self.offset--
	returnType := self.parseFieldType()
	if returnType == "" {
		self.causePanic()
	}
	self.parsed.returnType = returnType
}

func (self *MethodDescrptorParser) parseFieldType() string {
	switch self.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		self.offset--
		return ""
	}
}

func (self *MethodDescrptorParser) parseObjectType() string {
	str := self.methodDescriptor[self.offset:]
	index := strings.IndexRune(str, ';')
	if index == -1 {
		self.causePanic()
		return ""
	}
	st := self.offset - 1
	ed := self.offset + index + 1
	self.offset = ed
	return self.methodDescriptor[st:ed]
}

func (self *MethodDescrptorParser) parseArrayType() string {
	st := self.offset - 1
	self.parseFieldType()
	ed := self.offset
	return self.methodDescriptor[st:ed]
}

func (self *MethodDescrptorParser) checkStartParam() {
	if self.readUint8() != '(' {
		self.causePanic()
	}
}

func (self *MethodDescrptorParser) checkEndParam() {
	if self.readUint8() != ')' {
		self.causePanic()
	}
}

func (self *MethodDescrptorParser) checkFinish() {
	if self.offset != len(self.methodDescriptor) {
		self.causePanic()
	}
}

func (self *MethodDescrptorParser) readUint8() uint8 {
	b := self.methodDescriptor[self.offset]
	self.offset++
	return b
}

func (self *MethodDescrptorParser) causePanic() {
	panic("Bad descriptor: " + self.methodDescriptor)
}

package heap

import (
	"jvm/classfile"
	"strings"
)

/*
 * 方法区是运行时数据区的一块逻辑区域，由多个线程共享
 * 主要存放从class文件获取的类信息，使用结构体来表示类信息
 * accessFlags是类的访问标志
 * className, superClassName, interfaceNames代表类名、超类名以及接口名， 这些类名都是完全限定名，例如java/lang/String
 * initStarted字段表示类的<clinit>方法是否已经开始执行
 * jClass为java.lang.Class实例
 */
type Class struct {
	accessFlags       uint16
	className         string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	classLoader       *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.className = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) Name() string {
	return self.className
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

/*
 * 返回常量池
 */
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

/*
 * 返回静态变量表
 */
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

/*
 * 返回超类
 */
func (self *Class) SuperClass() *Class {
	return self.superClass
}

/*
 * 返回字段
 */
func (self *Class) Fields() []*Field {
	return self.fields
}

/*
 * 返回类加载器
 */
func (self *Class) Loader() *ClassLoader {
	return self.classLoader
}

/*
 * 返回java.lang.Class
 */
func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.className, "/", ".", -1)
}

/*
 * 返回方法
 */
func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) IsPublic() bool {
	return (self.accessFlags & ACC_PUBLIC) != 0
}

func (self *Class) IsFinal() bool {
	return (self.accessFlags & ACC_FINAL) != 0
}

func (self *Class) IsSuper() bool {
	return (self.accessFlags & ACC_SUPER) != 0
}

func (self *Class) IsInterface() bool {
	return (self.accessFlags & ACC_INTERFACE) != 0
}

func (self *Class) IsAbstract() bool {
	return (self.accessFlags & ACC_ABSTRACT) != 0
}

func (self *Class) IsSynthetic() bool {
	return (self.accessFlags & ACC_SYNTHETIC) != 0
}

func (self *Class) IsAnnotation() bool {
	return (self.accessFlags & ACC_ANNOTATION) != 0
}

func (self *Class) IsEnum() bool {
	return (self.accessFlags & ACC_ENUM) != 0
}

func (self *Class) isJlObject() bool {
	return self.className == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.className == "java/lang/Cloneable"
}

func (self *Class) isJioSerializable() bool {
	return self.className == "java/lang/Serializable"
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.className]
	return ok
}

/*
 * other类想要操作self类，需要满足两个条件：self类是public或者在同一个运行时包内
 */
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.className, "/"); i >= 0 {
		return self.className[:i]
	}
	return ""
}

/*
 * public static void main(String[] args)
 */
func (self *Class) GetMainMethod() *Method {
	return self.getMethod("main", "([Ljava/lang/String;)V", true)
}

/*
 * 返回类初始化方法
 */
func (self *Class) GetClinitMethod() *Method {
	return self.getMethod("<clinit>", "()V", true)
}

func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic && method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.className)
	return self.classLoader.LoadClass(arrayClassName)
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotID)
}

func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotID, ref)
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

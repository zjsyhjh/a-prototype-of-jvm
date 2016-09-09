package heap

import (
	"fmt"
	"jvm/classfile"
	"jvm/classpath"
)

/*
 * 类加载器, 依赖classpath来搜索和读取class文件
 * classMap字段记录已经加载的类的数据, key是类的完全限定名
 * 类加载的大致步骤: 找到class文件并读入内存；解析class文件，生成虚拟机可以使用的类数据，并放入方法区; 最后进行链接
 */
type ClassLoader struct {
	classPath   *classpath.ClassPath
	classMap    map[string]*Class
	verboseFlag bool
}

func NewClassLoader(classPath *classpath.ClassPath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		classPath:   classPath,
		classMap:    make(map[string]*Class),
		verboseFlag: verboseFlag,
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

/*
 * 先加载java.lang.Class类
 * 将已经加载好的类关联类对象
 */
func (self *ClassLoader) loadBasicClasses() {
	jClassClass := self.LoadClass("java/lang/Class")
	for _, class := range self.classMap {
		class.jClass = jClassClass.NewObject()
		class.jClass.extra = class
	}
}

/*
 * 加载基本类型类, void, int, double等，基本类型没有超类，也实现任何接口
 */
func (self *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		self.loadPrimitiveClass(primitiveType)
	}
}
func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		className:   className,
		classLoader: self,
		initStarted: true,
	}
	class.jClass = self.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	self.classMap[className] = class
}

/*
 * 加载类，先查找是否已经加载，若是，则返回类
 * 如果是数组类，则加载数组类，否则就加载普通类
 * 看java.lang.Class类是否已经加载，如果是，则给类关联类对象
 */
func (self *ClassLoader) LoadClass(className string) *Class {
	if class, ok := self.classMap[className]; ok {
		return class
	}
	var class *Class
	if className[0] == '[' {
		class = self.loadArrayClass(className)
	} else {
		class = self.loadNonArrayClass(className)
	}

	if jClassClass, ok := self.classMap["java/lang/Class"]; ok {
		class.jClass = jClassClass.NewObject()
		class.jClass.extra = class
	}
	return class
}

/*
 * 加载数组类
 */
func (self *ClassLoader) loadArrayClass(className string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		className:   className,
		classLoader: self,
		initStarted: true,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[className] = class
	return class
}

/*
 * 加载普通类(非数组类)
 */
func (self *ClassLoader) loadNonArrayClass(className string) *Class {
	data, entry := self.readClass(className)
	class := self.defineClass(data)
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", className, entry.String())
	}
	return class
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.classLoader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.className] = class
	return class
}

/*
 * readClass调用ClassPath的ReadClass方法
 */
func (self *ClassLoader) readClass(className string) ([]byte, classpath.Entry) {
	data, entry, err := self.classPath.ReadClass(className)
	if err != nil {
		panic("java.lang.ClassNotFoundException : " + className)
	}
	return data, entry
}

/*
 * 把class文件数据转换为class结构体
 * 调用classfile包下的parseClassFile()方法
 */
func parseClass(data []byte) *Class {
	cf, err := classfile.ParseClassFile(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

/*
 * 递归调用加载超类，除了java.lang.Object类
 */
func resolveSuperClass(class *Class) {
	if class.className != "java/lang/Object" {
		//fmt.Printf("className : " + class.className + " superClassName : " + class.superClassName + "\n")
		class.superClass = class.classLoader.LoadClass(class.superClassName)
	}
}

/*
 * 递归调用加载类的每一个接口
 */
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			/*
				fmt.Printf("ClassName : " + class.className)
				fmt.Printf(" InterfaceName : " + interfaceName)
				fmt.Println()
			*/
			class.interfaces[i] = class.classLoader.LoadClass(interfaceName)
		}
	}
}

/*
 * 类的链接分为两部分，验证和准备阶段
 * 有专门的验证算法，暂时不实现
 */
func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//todo
}

/*
 * 准备阶段主要给类变量分配空间并赋予初始值
 */
func prepare(class *Class) {
	calInstanceFieldSlotID(class)
	calStaticFieldSlotID(class)
	allocAndInitStaticVars(class)
}

/*
 * 计算实例字段的个数
 */
func calInstanceFieldSlotID(class *Class) {
	slotID := uint(0)
	if class.superClass != nil {
		slotID = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotID = slotID
			slotID++
			if field.isLongOrDouble() {
				slotID++
			}
		}
	}
	class.instanceSlotCount = slotID
}

/*
 * 计算静态字段的个数
 */
func calStaticFieldSlotID(class *Class) {
	slotID := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotID = slotID
			slotID++
			if field.isLongOrDouble() {
				slotID++
			}
		}
	}
	class.staticSlotCount = slotID
}

/*
 * 给类变量复制并分配空间
 */
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVars(class, field)
		}
	}
}

/*
 * 从常量池中加载常量池，然后给静态变量赋值
 */
func initStaticFinalVars(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstantValueIndex()
	slotID := field.SlotID()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			value := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotID, value)
		case "J":
			value := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotID, value)
		case "F":
			value := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotID, value)
		case "D":
			value := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotID, value)
		case "Ljava/lang/String;":
			fmt.Println("Ljava/lang/String;")
			goStr := cp.GetConstant(cpIndex).(string)
			javaStr := JString(class.Loader(), goStr)
			vars.SetRef(slotID, javaStr)
		}
	}
}

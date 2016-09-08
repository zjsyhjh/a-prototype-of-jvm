package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"float":   "F",
	"double":  "D",
	"char":    "C",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

/*
 * 如果是数组类名，描述符就是类名，直接返回
 * 如果是基本类型名，返回对应的是类型描述符
 * 否则就是普通的类型，则返回普通类名的形式
 */
func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}

	if d, ok := primitiveTypes[className]; ok {
		return d
	}

	return "L" + className + ";"
}

func getComponentClassName(className string) string {
	if className[0] == '[' {
		return toClassName(className[1:])
	}
	panic("Not array : " + className)
}

/*
 * 如果类型描述符以[开头，则是数组, 描述符即类名
 * 如果类型描述符以L开头，则是类描述符，去掉开头的L以及末尾的;就是类名
 */
func toClassName(descriptor string) string {
	/*
	 * 是数组类型
	 */
	if descriptor[0] == '[' {
		return descriptor
	}
	/*
	 * 是引用类型, 如Ljava/lang/String;
	 */
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}

	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}

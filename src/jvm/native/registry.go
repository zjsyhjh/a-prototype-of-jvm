package native

import "jvm/rtda"

/*
 * 本地方法的注册与查找
 */
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

/*
 * 类名+方法名+方法描述符确定key
 */
func Register(className, methodName, methodDescriptor string, nativeMethod NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = nativeMethod
}

/*
 * 根据类名+方法名+方法描述符返回本地方法
 */
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if nativeMethod, ok := registry[key]; ok {
		return nativeMethod
	}

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	//do nothing
}

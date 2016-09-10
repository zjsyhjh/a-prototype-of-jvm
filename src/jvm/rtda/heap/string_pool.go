package heap

import "unicode/utf16"

/*
 * java.lang.String是不可变的，JVM维护了一个字符串池用于节省内存
 */
/*
 * 使用map来保存字符串，key是Go字符串，value是Java字符串
 * Go字符串格式为utf8格式，Java字符串格式为utf16
 * 先转化成utf32，然后通过utf16的Encode包转化成utf16
 */
var internedStrings = map[string]*Object{}

func JString(classLoader *ClassLoader, goString string) *Object {
	if internedStr, ok := internedStrings[goString]; ok {
		return internedStr
	}

	chars := stringToUtf16(goString)
	jChars := &Object{classLoader.LoadClass("[C"), chars, nil}

	jStr := classLoader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goString] = jStr
	return jStr
}

/*
 * utf8 -> utf16
 */
func stringToUtf16(str string) []uint16 {
	runes := []rune(str)
	return utf16.Encode(runes)
}

/*
 * java.lang.String -> go string
 */
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

/*
 * utf16 -> utf8
 */
func utf16ToString(str []uint16) string {
	runes := utf16.Decode(str)
	return string(runes)
}

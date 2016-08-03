package classfile

/*
 * Deprecated和Synthetic是最简单的两种属性，仅起标记作用，不包含任何数据
 * Deprecated属性用于指出类、接口、字段或者方法已经不建议使用
 Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
 }
 * Synthetic属性用于标记源文件中不存在、由编译器生成的类成员
 Synthetic_attribute {
     u2 attribute_name_index;
     u4 attribute_length;
 }
*/
type MarkerAttribute struct{}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

/*
 * Deprecated和Synthetic属性都没有数据，因此方法为空
 */
func (ma *MarkerAttribute) readInfo(cr *ClassReader) {
	// do nothing
}

package classpath

import (
	"os"
	"strings"
)

/*
 * 路径分隔符，Linux下为':', windows下为';'
 */
const pathListSeparator string = string(os.PathListSeparator)

/*
 * 入口接口, 某个数据类型实现了该接口的所有方法，则实现了该接口
 * className :java/lang/Object.class
 */
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/*
 *
 */
func newEntry(path string) Entry {

	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}

package classpath

import (
	"errors"
	"strings"
)

/*
 * 保存多个Entry
 */
type CompositeEntry struct {
	compositeEntry []Entry
}

func newCompositeEntry(path string) *CompositeEntry {

	compositeEntry := &CompositeEntry{}

	for _, subpath := range strings.Split(path, pathListSeparator) {
		entry := newEntry(subpath)
		compositeEntry.compositeEntry = append(compositeEntry.compositeEntry, entry)
	}
	return compositeEntry
}

/*
 * 读取到名为className的类就返回
 */
func (cn *CompositeEntry) readClass(className string) ([]byte, Entry, error) {

	for _, entry := range cn.compositeEntry {
		bytes, classEntry, err := entry.readClass(className)
		if err == nil {
			return bytes, classEntry, nil
		}
	}

	return nil, nil, errors.New("Class not found: " + className)
}

func (cn *CompositeEntry) String() string {

	entrystr := make([]string, len(cn.compositeEntry))

	for index, entry := range cn.compositeEntry {
		entrystr[index] = entry.String()
	}

	return strings.Join(entrystr, pathListSeparator)
}

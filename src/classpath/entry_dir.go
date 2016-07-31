package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/*
 * 格式为：/pathto/xxx.class
 */
type DirEntry struct {
	absPath string
}

/*
 *先把路径转化为绝对路径，后返回一个DirEntry结构
 */
func newDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absPath}
}

func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {

	fileName := filepath.Join(de.absPath, className)
	bytes, err := ioutil.ReadFile(fileName)

	return bytes, de, err
}

func (de *DirEntry) String() string {
	return de.absPath
}

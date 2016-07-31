package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/*
 *Zip or Jar格式
 */
type ZipEntry struct {
	absPath string
	rc      *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {

	zipPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{zipPath, nil}
}

func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	if ze.rc == nil {
		var err = ze.openZip()
		if err != nil {
			return nil, nil, err
		}
	}

	var classFile = ze.findClassFile(className)

	if classFile == nil {
		return nil, nil, errors.New("Class not found: " + className)
	}

	bytes, err := readClassFile(classFile)

	return bytes, ze, err
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}

/*
 * 打开zip压缩包
 */
func (ze *ZipEntry) openZip() error {
	rc, err := zip.OpenReader(ze.absPath)
	if err == nil {
		ze.rc = rc
	}
	return err
}

/*
 *查找zip压缩包中的符合名称的类文件
 */
func (ze *ZipEntry) findClassFile(className string) *zip.File {
	for _, f := range ze.rc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

/*
 *读取某个类所在文件的内容
 */
func readClassFile(classFile *zip.File) ([]byte, error) {
	r, err := classFile.Open()
	defer r.Close()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

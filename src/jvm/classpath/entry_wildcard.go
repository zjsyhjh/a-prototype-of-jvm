package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/*
 * path参数：/pathto/*
 */
func newWildcardEntry(path string) *CompositeEntry {

	baseDir := path[:len(path)-1]
	compositeEntry := &CompositeEntry{}

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		/*
		 * 不遍历子目录
		 */
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			entry := newZipEntry(path)
			compositeEntry.compositeEntry = append(compositeEntry.compositeEntry, entry)
		}

		return nil
	}

	filepath.Walk(baseDir, walkFunc)

	return compositeEntry
}

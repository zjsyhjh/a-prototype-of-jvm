package classpath

import (
	"os"
	"path/filepath"
)

/*
 *ClassPath存放三类路径， 分别为启动类路径、扩展类路径以及用户类路径
 */
type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

/*
 * 构造ClassPath类型，赋予启动类路径、扩展类路径以及用户类路径值
 */
func ParseClassPathOption(jreOp, cpOp string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOp)
	cp.parseUserClassPath(cpOp)
	return cp
}

/*
 * 解析启动类路径和扩展类路径
 */
func (cp *ClassPath) parseBootAndExtClassPath(jreOp string) {
	jreDir := getJreDir(jreOp)

	// pathTo/jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClassPath = newWildcardEntry(jreLibPath)

	// pathTo/jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClassPath = newWildcardEntry(jreExtPath)
}

/*
 * 解析用户类路径，如果没有-cp选项，则从当前目录下找，否则从-cp指定的路径下寻找
 */
func (cp *ClassPath) parseUserClassPath(cpOp string) {
	if cpOp == "" {
		cpOp = "./"
	}
	cp.userClassPath = newEntry(cpOp)
}

/*
 * 获取JRE目录, 优先使用-Xjre选项查找，如果没有该选项，则在当前目录下找
 * 如果当前目录下也不存在JRE目录，则使用JAVA_HOME环境变量
 */
func getJreDir(jreOp string) string {
	if jreOp != "" && exists(jreOp) {
		return jreOp
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}

	panic("Coundn't find any JRE folder!")
}

/*
 * 检查文件或者目录是否存在
 * 存在返回true，否则flase
 */
func exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

/*
 *给输入的类名添加后缀.class, 并且读取指定类
 */
func (cp *ClassPath) readClass(className string) ([]byte, Entry, error) {
	className += ".class"

	if bytes, entry, err := cp.bootClassPath.readClass(className); err == nil {
		return bytes, entry, nil
	}

	if bytes, entry, err := cp.extClassPath.readClass(className); err == nil {
		return bytes, entry, nil
	}
	return cp.userClassPath.readClass(className)
}

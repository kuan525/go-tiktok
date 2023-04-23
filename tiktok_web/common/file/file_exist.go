package file

import "os"

// Exists 判断文件或目录是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

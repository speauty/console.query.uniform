package util

import (
	"os"
	"strings"
)

// CheckDirExist 检测目录是否存在
func CheckDirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建目录
func CreateDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// CreateDirRecursion 递归创建目录
func CreateDirRecursion(path string) error {
	pathArr := strings.Split(path, "/")
	currentPath := ""
	for _, tmpPath := range pathArr {
		currentPath += tmpPath + "/"
		if strings.Contains(tmpPath, ".") {
			continue
		}
		exist, err := CheckDirExist(currentPath)
		if err != nil {
			return err
		}
		if exist == false {

			if err = CreateDir(currentPath); err != nil {
				return err
			}
		}
	}
	return nil
}

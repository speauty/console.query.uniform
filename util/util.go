package util

import (
	"console.query.uniform/kernel/constants"
	"io/ioutil"
	"os"
	"runtime"
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

// LoadFile2Str 直接加载文件
func LoadFile2Str(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	return string(bytes), err
}

// GetOS 获取操作系统
func GetOS() string {
	return runtime.GOOS
}

// GetArchitecture 获取架构
func GetArchitecture() string {
	return runtime.GOARCH
}

// GetDS 获取分隔符
func GetDS() string {
	if GetOS() == constants.SysOsWindows {
		return "\\"
	} else {
		return "/"
	}
}

// GetBinName 获取可执行文件名称
func GetBinName() string {
	bin := strings.Split(os.Args[0], GetDS())
	return bin[len(bin)-1]
}

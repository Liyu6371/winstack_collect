package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// MakeFile 创建文件，如果存在则不继续创建
func MakeFile(dir, file string) error {
	p := filepath.Join(dir, file)
	exist, err := FileExists(p)
	if err != nil {
		return fmt.Errorf("unable to check the file is exists or not, error: %s", err)
	}
	// 文件存在的情况直接返回
	if exist {
		return nil
	}
	// 开始创建文件
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to make dir, error: %s", err)
	}
	f, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("create file error: %s", err)
	}
	defer f.Close()
	// 创建成功
	return nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	// 文件不存在的情况
	if os.IsNotExist(err) {
		return false, nil
	}
	// 其他的错误情况，返回 false 以及错误信息
	if err != nil {
		return false, err
	}
	// 文件存在的情况
	return true, nil
}

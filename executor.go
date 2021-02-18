package gutils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ExecFolder 获取执行文件所在的目录，
// 与os.Getwd() 区别是
//	os.Getwd() 不包含程序的目录
// 	本方法包含目录
func ExecFolder() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	re, err := filepath.Abs(file)
	if err != nil {
		fmt.Printf("The eacePath failed: %s\n", err.Error())
	}
	fmt.Print("The path is ", re)
	return filepath.Abs(file)
}
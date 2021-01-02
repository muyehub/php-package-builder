package builder

import (
	"fmt"
	"io"
	"os"
)

const (
	// PACKAGEDIR 包目录常量
	PACKAGEDIR string = "./"
)

// Build 创建文件夹
func Build(name, email string) {
	CreateDir(PACKAGEDIR + "bin/")
	CreateDir(PACKAGEDIR + "build/")
	CreateDir(PACKAGEDIR + "docs/")
	CreateDir(PACKAGEDIR + "config/")
	CreateDir(PACKAGEDIR + "src/")
	CreateDir(PACKAGEDIR + "tests/")
	CreateDir(PACKAGEDIR + "vendor/")
}

// CreateDir 创建文件夹
func CreateDir(dirname string) {
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		fmt.Println(err)
	}
	CreateFile(dirname + ".gitkeep")
}

// CreateFile 创建文件
func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}

// CopyFile 拷贝文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

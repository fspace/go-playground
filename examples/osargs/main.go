package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
		return
	}
	root, err := filepath.Abs(os.Args[1]) // 获取绝对路径
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}
	fmt.Println("Listing files in ", root)

	var c struct {
		files int
		dirs  int
	}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// 遍历文件树统计文件和目录的数量
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})

	fmt.Printf("Total: %d files in %d diretories", c.files, c.dirs)

}

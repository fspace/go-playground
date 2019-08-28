package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"io"
	"os"
)

func removeDemo1() {
	if err := os.Remove("file.txt"); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func createDemo() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
}
func renameDemo() {
	if err := os.Rename("file.txt", "file2.txt"); err != nil {
		fmt.Println("Rename Err:", err)
	}
}

func copyFile(from, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	createDemo()
	// removeDemo1()

	renameDemo()

	n, err := copyFile("file2.txt", "file3.txt")
	if err != nil {
		fmt.Println("copyFile Error:", err)
	}
	_ = n
}

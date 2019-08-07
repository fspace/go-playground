package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readingAllAtOnce() {
	// 使用 512 字节缓存
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
		return
	}
	fmt.Println("os args: ", os.Args)
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(b))
}

func main() {
	readingAllAtOnce()
}

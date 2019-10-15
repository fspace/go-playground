package main

import (
	"fmt"
	"os"
)

// 退出值只能是8 bit 值
func main() {
	fmt.Println("Hello, Playground!")
	defer fmt.Println("Hello, playground") // this will not be called

	os.Exit(-1) // -1 % 256 = 255 // ！！！ 是不是改版了！
}

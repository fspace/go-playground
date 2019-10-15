package main

import "fmt"

// 退出码是2
func main() {
	defer fmt.Println("Hello, Playground !")
	panic("panic")
}

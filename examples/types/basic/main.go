package main

import "fmt"

// 8 16 32 64 范围问题
// 超范围后 编译期都会报错的
func intDemo() {
	var i1 int8 = -128
	fmt.Println("i1 is :", i1)

	//var i2 int8 = -129
	//fmt.Println("i2 is :", i2)

	var ui uint8 = 255
	fmt.Println("u int :", ui)
}

func main() {
	intDemo()
}

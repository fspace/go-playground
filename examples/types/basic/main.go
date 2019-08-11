package main

import "fmt"

// -1 问题       首位 0 1 表正负如：  0000000000000000     1000000000000...  代表 +0  -0 两者其实可以去掉一个 所以范围就-1了
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

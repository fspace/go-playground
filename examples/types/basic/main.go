package main

import (
	"fmt"
	"unsafe"
)

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

// byte 0-255 类似其他语言的 char类型
// rune int32 别名类型 可以取负 这个比较以外！

func intDemo2() {
	var a = 40
	fmt.Printf("a type is %T", a)
	// 类型 及其占用的字节数
	fmt.Printf(" 大小是 %d ", unsafe.Sizeof(a))
}

func sizeDemo() {
	a := struct {
		A int
		B float32
	}{A: 20, B: 22.2}
	fmt.Printf("type a is %T and size is %d ", a, unsafe.Sizeof(a))
}

func main() {
	intDemo()
	intDemo2()
	sizeDemo()

}

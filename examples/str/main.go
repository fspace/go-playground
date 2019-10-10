package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func t1() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello, world)

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println(s1, s2)
	// 下面仅仅是演示 非推荐做法
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
}

func t2() {
	fmt.Printf("%#v\n", []byte("Hello, 世界"))

	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界
	// 部分损坏
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �abc

	// 迭代损坏
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
	/*
		// 0 65533 // \uFFFD, 对应 �
		// 1 0 // 空字符
		// 2 0 // 空字符
		// 3 30028 // 界
		// 6 97 // a
		// 7 98 // b
		// 8 99 // c
	*/
	// 遍历
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	// 或者是采用传统的下标方式遍历字符串的字节数组：
	const s = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %x\n", i, s[i])
	}
}

func t3() {
	fmt.Printf("%#v\n", []rune("世界"))
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
}

func main() {
	//t1()
	//t2()
	//t3()
	t5()
}

// rune 与 len 统计数的区别
func t5() {
	s := "hello世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCount([]byte(s)))
}

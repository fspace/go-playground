package main

import (
	"fmt"
	"unsafe"
)

func boolDemo() {
	b := true
	fmt.Println("b is ", b)

	fmt.Printf("type of b is %T\n ", b)
	fmt.Printf("size of b is %d \n ", unsafe.Sizeof(b))
}

func main() {
	boolDemo()

	boolDemo2()

	// 零值
	zeroVal()
}

func boolDemo2() {
	var n bool = true
	fmt.Printf("%v , %T\n", n, n)

	m := 1 == 2
	fmt.Printf("%v , %T\n", m, m)
}

func zeroVal() {
	var n bool
	fmt.Printf("%v , %T\n", n, n)
}

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

	operators()
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

func operators() {
	// bool 型可用的逻辑操作符 （与或非）
	var isTrue bool = true
	var isFalse bool = false
	// AND
	if isTrue && isFalse {
		fmt.Println("Both Conditions need to be True")
	}
	// OR - not exclusive
	if isTrue || isFalse {
		fmt.Println("Only one condition needs to be True")
	}

	if !isFalse {
		fmt.Println("is not false means: ", !isFalse)
	}
}

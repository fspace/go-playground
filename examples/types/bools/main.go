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
}

package main

import (
	"fmt"
	"unsafe"
)

type Null struct {
}

type P1 struct {
	A string
	B Null
}

type P2 struct {
	B Null
	A string
}

func main() {
	fmt.Println("P1 sizeOf ", unsafe.Sizeof(P1{}),
		"AlignOf ", unsafe.Alignof(P1{}),
		"OffsetOf A ", unsafe.Offsetof(P1{}.A),
		"OffsetOf B ", unsafe.Offsetof(P1{}.B))

}

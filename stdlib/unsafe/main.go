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

// henrylee2cn:
// 两个结构体因为Null字段位置不同，而size相差8。我觉得应该是对齐的问题。单独的Null结构体size为0，align为1；
// 在P1和P2中作为字段，如果是放在前面 align为1，会被对齐为8，放在最后则不需要再对齐了。
func main() {
	fmt.Println("P1 sizeOf ", unsafe.Sizeof(P1{}),
		"AlignOf ", unsafe.Alignof(P1{}),
		"OffsetOf A ", unsafe.Offsetof(P1{}.A),
		"OffsetOf B ", unsafe.Offsetof(P1{}.B))

}

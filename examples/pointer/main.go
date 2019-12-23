package main

import "fmt"

func main() {
	basic2()
}

func basic() {
	a := 42
	fmt.Println(a)
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
}
func basic2() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)

	fmt.Println(&a, b)
	// dereference
	fmt.Println(a, *b)

	// 修改
	a = 27
	fmt.Println(a, *b)
	fmt.Printf("addr: %p", b)
}

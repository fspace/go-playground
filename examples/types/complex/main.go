package main

import "fmt"

func main() {
	demo()
	demo1()

	ops()
	selectorFunc()

	zeroVal()
}
func demo() {
	var n complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n ", n, n)
}
func demo1() {
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", n, n)
}

func selectorFunc() {
	fmt.Println(">--selector-func------------")
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n ", real(n), real(n)) // 实部
	fmt.Printf("%v, %T\n ", imag(n), imag(n)) // 虚部
	fmt.Println("======================================")

	var n2 complex128 = 1 + 2i
	fmt.Printf("%v, %T\n ", real(n2), real(n2)) // 实部
	fmt.Printf("%v, %T\n ", imag(n2), imag(n2)) // 虚部
}

func ops() {
	a := 1 + 2i
	b := 2 + 5.2i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

func zeroVal() {
	fmt.Println(">-- zero-val------------------")
	var a complex64
	fmt.Printf("%v , %T \n", a, a)
}

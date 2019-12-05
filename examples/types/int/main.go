package main

import "fmt"

func main() {
	intDemo()

	intOps()
	bitOps()

	zeroVal()
}

func intDemo() {
	n := 42
	fmt.Printf("%v, %T", n, n) // int 型长度依据所运行的平台长度不定 比如32位|64位
}

func intOps() {
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
func bitOps() {
	fmt.Println(">--bit-ops--------------------")
	a := 10            // 1010
	b := 3             // 0011
	fmt.Println(a & b) //
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	// shifting
	a = 8               // 2 ^ 3
	fmt.Println(a << 3) // 2^3 * 2^3
	fmt.Println(a >> 3) // 2^3 / 2^3
}

func zeroVal() {
	var n int
	fmt.Printf("%v, %T \n", n, n)
}

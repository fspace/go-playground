package main

import "fmt"

func main() {
	intDemo()

	intOps()
	bitOps()

	zeroVal()

	ConversionOfTypes()
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

// 边界溢出
func overflowsBboundary() {
	fmt.Println("Hello World")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint) // prints out -127
}

func ConversionOfTypes() {
	var men uint8
	men = 5
	var women int16
	women = 6

	//var people int
	//// this throws a compile error
	//people = men + women
	// this handles converting to a standard format
	// and is legal within your go programs
	people := int(men) + int(women)

	fmt.Println("people: ", people)
}

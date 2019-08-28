package main

import (
	"fmt"
	"strconv"
)

// see studygolang.com/pkgdoc
func toStrDemo1() {
	var n1 int = 99
	var n2 float64 = 123.456
	var b bool = true
	var myChar byte = 'b'
	var str string

	str = fmt.Sprintf("%d", n1)
	fmt.Println(str)

	str = fmt.Sprintf("%f", n2)
	fmt.Println(str)

	//str = fmt.Sprintf("%b",b)
	str = fmt.Sprintf("%t", b)
	fmt.Println(str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Println(str)
}

func toStrDemo2() {
	var n1 int = 99
	var n2 float64 = 123.456
	var b bool = true
	//var myChar byte = 'b'
	var str string
	str = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("str type is %T and value is %q\n ", str, str)

	str = strconv.FormatFloat(n2, 'f', 10, 64)
	fmt.Println(str)

	str = strconv.FormatBool(b)
	fmt.Println(str)

}

func main() {
	toStrDemo1()
	toStrDemo2()
}

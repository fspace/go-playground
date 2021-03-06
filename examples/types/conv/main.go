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

	typeConvert()
}

func typeConvert() {
	var i float32 = 42.5
	fmt.Printf("%v , %T \n ", i, i)

	var j int
	j = int(i) /// 截断
	fmt.Printf("%v, %T \n", j, j)

	var s string
	s = string(j) // int 转string！  这个转为42所代表的utf8码点了
	fmt.Printf("%v, %T \n", s, s)

	s = strconv.Itoa(j) // 使用包转为字面量
	fmt.Printf("%v, %T \n", s, s)
}

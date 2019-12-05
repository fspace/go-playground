package main

import "fmt"

/**
## utf-8 字符编码序列
*/

func strDemo1() {
	addr := "中国 长城 \n hello china"
	fmt.Println(addr)

	str2 := `
	你好呀 golang
	我是你爸爸

code:
package main

import "fmt"

func main(){ 
...
}
   
`
	fmt.Println(str2)
}

func strDemo2() {
	s := "hello " + "world "
	s += "hahaha !"
	fmt.Println(s)

	s2 := "hello " + "hello " + "hello " + "hello " + "hello " +
		"hello " +
		"hello " +
		"hello " +
		"end !"
	fmt.Println(s2)

}

func main() {
	//strDemo1()
	//strDemo2()
	strDemo()
	strOps()

	conv()

	zeroVal()
}

func strDemo() {
	s := "this is a string"
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", s[2], s[2])         // byte类型  unit8 是别名
	fmt.Printf("%v, %T \n", string(s[2]), s[2]) // 转型
}

func strOps() {
	s := "this is a string"
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
}

func conv() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

func zeroVal() {
	var s string
	fmt.Printf("%v, %T\n", s, s)
}

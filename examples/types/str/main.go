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
	strDemo1()
	strDemo2()
}

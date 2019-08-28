package main

import "fmt"

/**
	golang 无char类型 用byte
	## 字符的本质是一个整数
    ## aciII 0~255 之间的用byte就够了  超过的就需要其他int型  或者rune（别名 int32）
	## 网络上可以搜索：ASCII 码表   unicode码表|utf8 码表
*/

func charDemo1() {
	var c1 byte = 'A'
	fmt.Println("c1 is :", c1)

	var c2 byte = '0'
	// 数字类型输出
	fmt.Println("c2 = ", c2)
	// 字符类型输出
	fmt.Printf("c1 = %c c2 = %c \n", c1, c2)

	// var c3 byte = '北' //  constant 21271 overflows byte
	var c3 = '北'
	fmt.Printf("c3 is %c and the type is %T , the code value is : %d \n", c3, c3, c3)
}

func charDemo2() {
	c := '\n'
	fmt.Printf("%c is , code value is %d \n", c, c)

	c2 := '中'
	fmt.Println("c2 is :", c2)
	c3 := 20013
	fmt.Printf("c3 is %c \n", c3)

	// NOTE 字符做整数运算
	c_a := 'a'
	c_A := 'A'
	fmt.Println("A - a is:", c_A-c_a)
	cB := 'b' - 32
	fmt.Printf("b is : %c \n ", cB)
}

func main() {
	charDemo1()

	charDemo2()
}

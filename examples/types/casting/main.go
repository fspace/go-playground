package main

import "fmt"

/**
## 必须强转
## 被转换者 本身类型不可变  相当于函数式中的参数
*/
func demo1() {
	i := 100
	var n float64 = float64(i)
	fmt.Println("n is ", n)

	var i2 int8 = int8(i)
	fmt.Printf("i2 is %v \n", i2)
	var i3 int64 = int64(i)
	fmt.Printf("i3 is %v \n", i3)

	fmt.Printf("i type is %T\n ", i)
}

func demo2() {
	var n1 int64 = 666666666
	var n2 int8 = int8(n1)
	fmt.Printf("n2 is %v \n", n2)
	fmt.Printf("bin n1 is %b \n", n1)
	fmt.Printf("bin n2 is %b \n", n2)
}

func main() {
	demo1()

	demo2()
}

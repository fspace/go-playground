package main

import "fmt"

/**
  - 单精度 float32
  - 双精度  float64
  三个部分  符号位 + 指数位 + 尾数位
  尾数部分可能丢失 造成精度损失

	## 开发中 如无特殊需要 推荐使用float64
	## golang float类型 是操作系统无关的
*/
func floatDemo() {
	var price = 100.05
	fmt.Println("price is :", price)
	fmt.Printf("type : %T \n", price) // 默认浮点类型是float64哦

	var n1 float32 = 0.001
	var n2 float64 = 22333333333.33333
	fmt.Println("number1 is ", n1, "number2 is ", n2)

	var f1 float32 = -123.0000901 // 末位损失掉了
	var f2 float64 = -123.0000901
	fmt.Println("f1 is :", f1, "f2 is :", f2)

	n3 := 5.12
	n4 := .12
	fmt.Println("n3 IS ", n3, "n4: ", n4)
	// 科学计数法
	n5 := 5.1234e2
	fmt.Println("n5 is :", n5)
	n6 := 5.1234E2
	fmt.Println("n6 is :", n6)
	n7 := 5.00E-2
	fmt.Println("n7 is :", n7)
}

func main() {
	//floatDemo()
	fltDemo2()
	fltOps()

	maxFloat()

	converting()
}

func fltDemo2() {
	n := 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("%v, %T \n", n, n)
}

func fltOps() {
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
func fltOps2() {
	//a := 10.2
	//b := 3.7
	//fmt.Println(math.)
}

// 边界值
func maxFloat() {
	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 16777216
	fmt.Println(maxFloat32 + 2000000) // 16777216
}

// 类型转换
func converting() {
	// converting from int to float
	var myint int
	myfloat := float64(myint)
	fmt.Println("myfloat: ", myfloat)

	// converting from float to int
	var myfloat2 float64
	myint2 := int(myfloat2)
	fmt.Println("myint2: ", myint2)
}

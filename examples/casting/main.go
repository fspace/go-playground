package main

import "fmt"

// 两类转型           接口《--》其他类型            具体类型转型 | 数字类转型（涉及精度损失)

// With concrete types, casting can happen between types that have the same memory structure, or it can happen between numerical types:
type N [2]int

var n = N{1, 2}

var m [2]int = [2]int(n)

var a = 3.14
var b int = int(a) //// numerical types can be casted, in this case a will be rounded to 3

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recover info:", r)
		}
	}()
	var i interface{} = "hello"
	x, ok := i.(int) // ok will be false
	if ok {
		fmt.Println("x is ", x)
	} else {
		fmt.Println("i is not a int val ")
	}
	typeSwitch()

	y := i.(int) // this will panic
	fmt.Println(y)
	z, ok := i.(string) // ok will be true
	fmt.Println(z)
}

func typeSwitch() {
	var a interface{} = 10
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	case string:
		fmt.Println("a is a string")
	}
}

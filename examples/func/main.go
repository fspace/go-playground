package main

import (
	"fmt"
	"playgo/examples/func/utils"
)

// 多个参数和多个返回值
func Swap(a, b int) (int, int)  {
	return b, a
}

// 可变数量的参数
// more 对应的 []int 切片类型
func Sum(a int , more ...int) int  {
	for _ ,v := range more{
		a += v
	}
	return a
}
// 可变参数
func myPrint(a ...interface{})  {
	fmt.Println(a...)
}

// 命名返回值
func Find(m map[int]int , key int) (value int , ok bool)  {
  value, ok = m[key]
 return
}

// defer
// 如果返回值命名了，可以通过名字来修改返回值，也可以通过 defer 语句 在 return 语句之后修改返回值：
// 。闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问。
func Inc()(v int)  {
	defer func() {v++}()

	return 42
}

func problemOfDefer()  {
	for i := 0 ; i<3; i++ {
		defer func() {println(i)}()
	}
}
// -------------------------------------------------------------------------------------------  \
//					for 内循环是不建议的 此段仅做示意
// repair the defer effect
func deferRepair()  {
	// FIXME 块作用域的变量 可以看做是形参！
	for i:=0; i<3; i++ {
		i := i // 定义一个循环体内部的变量i
		defer func() {println(i)}()
	}
}
func fixDefer2()  {
	println("fix the defer")
	for i:= 0; i<3; i++ {
		defer func(i int ) {
			println(i)
		}(i)
	}
}
// -------------------------------------------------------------------------------------------  /


func main() {
	fmt.Println(utils.Add(1,2))
	fmt.Println(utils.Add2(1,2))

	fmt.Println(Swap(1, 2))
	fmt.Println(Sum(1, 2 , 4, 5))

	var a = []interface{}{123, "abc"}
	myPrint(a...) // 解包
	myPrint(a)

	m := map[int]int{ 1:2 , 2:3}
	v , ok := Find(m,2)
	if ok {
		fmt.Println("we found it : " ,v)
	}

	fmt.Println(Inc())
	// defer的坑
	problemOfDefer()
	// 修复
	deferRepair()
	fixDefer2()
}

package main

import "fmt"

type MyIface interface {
}
type User struct {
	Name string
	Age  int
}

func Declaration() {
	// 0 值
	var a int
	println(a)

	var str string
	fmt.Printf("str: %v . it's val is equal to \"\" ? %v ", str, str == "")

	var b bool
	fmt.Printf("\n b is :%v \n", b)
	var u *User
	fmt.Printf("user is %v \n ", u)
	var i MyIface
	fmt.Printf("i is %v \n ", i)
	var s []int
	fmt.Printf("s is %v \n ", s)
	var m map[string]interface{}
	fmt.Printf("m is %v \n ", m)
	var c chan interface{}
	fmt.Printf("c is %v\n ", c)
}

func main() {
	Declaration()
}

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

func inferredType() {
	a1 := 1 // short declaration
	fmt.Printf("a type is :%T \n", a1)

	var a2 = 2
	fmt.Printf("a2 type is :%T \n", a2)

	//
	var d interface{} = 1
	_ = d
	//The initialization of a type using a built-in function is as follows:
	var a = new(*int)                   // pointer of a new in variable
	sliceEmpty := make([]int, 0)        // slice of int of size 0, and a capacity of 0
	sliceCap := make([]int, 0, 10)      // slice of int of size 0, and a capacity of 10
	var map1 = make(map[string]int)     // map with default capacity
	var map2 = make(map[string]int, 10) // map with a capacity of 10
	var ch1 = make(chan int)            // channel with no capacity (unbuffered)
	var ch2 = make(chan int, 10)        // channel with capacity of 10 (buffered)
	fmt.Println(
		a,
		sliceCap,
		sliceEmpty,
		map1,
		map2,
		ch1,
		ch2,
	)
}

func dclDemo() {
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	var name, age, sex = "yiqing", 18, 1
	fmt.Println("name\tage\t\tsex")
	fmt.Printf("%s\t%d\t%d \n", name, age, sex)

	// 类型推导
	nickname, address := "yiqing", "china xian"
	fmt.Println("name:", nickname, "addr:", address)
}

func main() {
	Declaration()
	inferredType()
	dclDemo()
}

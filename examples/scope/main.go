package main

import "fmt"
import "errors"

// 三个级别的scope  ： package   func   block

var myInt = 0

func myFunc() {
	fmt.Println("package main var , myInt :", myInt)

	var myInt = 2
	fmt.Println("func var myInt:", myInt)

	if true {
		var myInt = 3
		fmt.Println("block scope myInt:", myInt)

		//	var myVar2 = "my var 2 in block" // 块级变量 外部不可见
	}
	// print(myVar2)
	fmt.Println("now the func scoped var myInt is :", myInt)
}
func funcScope() {
	// this exists in the outside block
	var err error
	// this exists only in this block, shadows the outer err
	if err := errors.New("Doh!"); err != nil {
		fmt.Println(err) // this not is changing the outer err
	}

	fmt.Println(err) // outer err has not been changed
}

func main() {
	myFunc()
	fmt.Println("pkg var myInt IS:", myInt)

	funcScope()
}

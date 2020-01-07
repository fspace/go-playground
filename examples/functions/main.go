package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("struct", "demo for studying struct ")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("mr", "Multiple Results", cli.ActionCommand(MultipleResults))
	app.Command("af", "Anonymous Func: 匿名函数", cli.ActionCommand(AnonymousFunc))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
func myfunction(firstName string, lastName string) string {
	fullname := firstName + " " + lastName
	return fullname
}

func basics() {
	fmt.Println("Hello World")

	fullName := myfunction("Elliot", "Forbes")
	fmt.Println(fullName)
}

func MultipleResults() {
	fmt.Println("Hello World")

	// we can assign the results to multiple variables
	// by defining their names in a comma separated list
	// like so:
	fullName, err := myfunction2("Elliot", "Forbes")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)
}
func myfunction2(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func AnonymousFunc() {
	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5

	addTen := addN(10)
	fmt.Println(addTen(1))  //
	fmt.Println(addTen(2))  //
	fmt.Println(addN(2)(3)) // 2+3
}
func addOne() func() int {
	// 如果接受参数  那么该函数就可以维护一个类似 对象局部变量状态一样的功能了！柯里化（curry）
	var x int
	// we define and return an
	// anonymous function which in turn
	// returns an integer value
	return func() int {
		// this anonymous function
		// has access to the x variable
		// defined in the parent function
		x++
		return x + 1
	}
}
func addN(n int) func(int) int {
	// 相当于把函数 f1(a, b)  转换为形式： fx(a)(b)
	return func(a int) int {
		return n + a
	}
}

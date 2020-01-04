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
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basicSyntax))
	app.Command("params", "parameters demo", parametersCmd)
	app.Command("mp", "multiple params example", cli.ActionCommand(multipleParams))
	app.Command("pv", "pass values", cli.ActionCommand(passValues))
	app.Command("pp", "pass pointers", cli.ActionCommand(passByPointers))
	app.Command("vf", "Variadic functions", cli.ActionCommand(variadicParams))

	app.Command("rv", "return value", cli.ActionCommand(returnValues))
	app.Command("mrv", "multiple return values", cli.ActionCommand(multipleReturnValues))

	app.Command("af", "anonymous functions ", cli.ActionCommand(AnonymousFunctions))
	app.Command("m", " methods", cli.ActionCommand(methods))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
// AGENDA
// basicSyntax
func basicSyntax() {
	fmt.Println("this is a function ")
}

// parametersCmd
func parametersCmd(cmd *cli.Cmd) {
	msg := "hello go"

	msg2 := cmd.StringArg("MSG", "hi yiqing", "Msg passed to the function")

	cmd.Action = func() {
		sayMsg(msg)
		fmt.Println("msg2 from cli :", *msg2)
	}

}
func sayMsg(msg string) {
	fmt.Println(msg)
}
func multipleParams() {
	for i := 1; i < 5; i++ {
		sayMsg2("Hello go !", i)
	}
	sayGreeting("hello ", "yiqing")
}
func sayMsg2(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is :", idx)
}
func sayGreeting(greeting string, name string) {
	fmt.Println(greeting, name)
}
func sayGreeting2(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted" // change the param
	fmt.Println(name)
}
func sayGreeting3(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted" // change the param
	fmt.Println(*name)
}
func passValues() {
	greeting := "Hello"
	name := "Qing"
	sayGreeting2(greeting, name)
	fmt.Println(name)
}
func passByPointers() {
	greeting := "Hello"
	name := "Qing"
	sayGreeting3(&greeting, &name)
	fmt.Println(name)
}
func variadicParams() {
	// @see https://gobyexample.com/variadic-functions
	// @see https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085
	// @see https://golangbot.com/variadic-functions/
	// @see https://www.geeksforgeeks.org/variadic-functions-in-go/
	// @see https://medium.com/rungo/variadic-function-in-go-5d9b23f4c01a
	// @see https://www.digitalocean.com/community/tutorials/how-to-use-variadic-functions-in-go
	// @see https://www.golangprograms.com/go-language/variadic-functions.html
	// @see https://yourbasic.org/golang/variadic-function/  不错的go知识网站： https://yourbasic.org/golang/
	// @see https://www.golangprograms.com/pass-different-types-of-arguments-in-variadic-function.html
	// @see https://programming.guide/go/three-dots-ellipsis.html
	// @see http://blog.stoneriverelearning.com/a-definitive-guide-to-variadic-functions-in-golang/

	// @see https://objectcomputing.com/resources/publications/sett/january-2019-way-to-go-part-2
	// ## BOOKS
	// @see https://www.golang-book.com/books/intro/7
	// @see https://go101.org/article/function.html

	// @see https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
	// @see https://golang.org/ref/spec
	sum(1, 2, 3, 4, 5)
}
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is :", result)
}

// --------------------------------
// returnValues
func returnValues() {
	s := sum2(1, 2, 3, 4, 5)
	fmt.Println("The sum is :", s)

	s2 := sum3(1, 2, 3, 4, 5)
	fmt.Println("The sum 2 is :", *s2)

	s3 := sum4(1, 2, 3, 4, 5)
	fmt.Println("Named return value: ", s3)

	d := divide(5.0, 0.0)
	fmt.Println(d)
}
func multipleReturnValues() {
	d, err := divide2(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
func sum2(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	//fmt.Println("The sum is :", result)
	return result
}
func sum3(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	//fmt.Println("The sum is :", result)
	return &result
}
func sum4(values ...int) (result int) {
	fmt.Println(values)
	// result := 0
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) float64 {
	return a / b
}
func divide2(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func AnonymousFunctions() {
	func() {
		fmt.Println("hello Go!")
	}()

	var f func() = func() {
		fmt.Println("hello qing!")
	}
	f()
	// -------------
	var divide func(float64, float64) (float64, error)
	divide = func(a float64, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("err2: cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
func FuncAsTypes() {

}

// ----------------------------------------------------------
func methods() {
	g := greater{
		greeting: "Hello",
		name:     "Go",
	}
	g.great()
	fmt.Println("The new name is :", g.name)
	g.great2()
	fmt.Println("The new name is :", g.name)

}

type greater struct {
	greeting string
	name     string
}

func (g greater) great() {
	fmt.Println(g.greeting, g.name)
}
func (g *greater) great2() {
	// 用指针来修改成员变量
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

// ===========================================================

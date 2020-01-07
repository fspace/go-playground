package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"playgo/examples/method/lib"
)

type A int

func (a A) Foo() { fmt.Println("Foo method of type A , and a is :", int(a)) }

// ========================================================================================= |
type ErrKey string

func (e ErrKey) Error() string {
	//return fmt.Errorf("key %q not found", e)
	return fmt.Sprintf("key %q not found", e)
}

var _ error = ErrKey("") // 确保类型实现了接口  可以在编译期发现问题
// ========================================================================================= |

func main() {
	//a := A(1)
	//a.Foo()  // Call the method on an instance of the type
	//A.Foo(a) // Call the method on the type and passing an instance as argument
	//
	//ms := lib.MyStruct{}
	//ms.Foo()

	os.Exit(realMain())
}

func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==================================== ========================================================
// Methods typically act upon a given object, i.e. guitarist.Update(params) and using it in this fashion is typically
// far preferential than doing UpdateGuitarist(guitarist, params) when it comes to writing your code.
/**
func UpdateGuitarist(guitarist *Guitarist, params ParamsStruct) {
  fmt.Println("This is a simple function")
}

// Calling this function
UpdateGuitarist(guitarist, params)

//  等价的对象形式：
 func (g *Guitarist) Update(params ParamsStruct) {
  fmt.Println("This is a simple method")
}

// this is far nicer in my opinion
myGuitarist.Update(params)

*/
// ==================================== ========================================================
func basics() {
	var employee lib.Employee
	employee.Name = "Elliot"
	employee.UpdateName("Forbsey")
	employee.PrintName()
}

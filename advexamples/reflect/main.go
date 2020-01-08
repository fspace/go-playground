package main

// https://gfw.go101.org/article/reflection.html
// https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/
import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
)
import "reflect"

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("r2i", "reflectObj to interface", cli.ActionCommand(reflectObj2interface))
	app.Command("mv", "modify Value", cli.ActionCommand(modifyValue))
	app.Command("mv2", "modify Value the right way", cli.ActionCommand(modifyValueOK))

	app.Command("it", "Implements Type", cli.ActionCommand(ImplementsType))
	app.Command("if", "invoke Func", cli.ActionCommand(invokeFunc))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

type F func(string, int) bool

func (f F) Validate(s string) bool {
	return f(s, 32)
}

func demo0() {
	var x struct {
		n int
		f F
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())     // struct
	fmt.Println(tx.NumField()) // 2
	tf := tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1
	fmt.Println(tf.NumMethod())          // 1
	ts, ti, tb := tf.In(0), tf.In(1), tf.Out(0)
	fmt.Println(ts.Kind(), ti.Kind(), tb.Kind()) // string int bool
}

// ---------

func basics() {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
}

func reflectObj2interface() {
	v := reflect.ValueOf(1)
	i2 := v.Interface().(int)
	fmt.Printf("%T: %v\n", i2, i2)
}

func modifyValue() {
	i := 1
	v := reflect.ValueOf(i)
	v.SetInt(10)
	fmt.Println(i)
}

func modifyValueOK() {
	/**
	想要修改原有的变量我们只能通过如下所示的方法，首先通过 reflect.ValueOf 获取变量指针，然后通过 Elem 方法获取指针指向的变量并调用 SetInt 方法更新变量的值
	*/
	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
}

// ----------------
type CustomError struct{}

func (*CustomError) Error() string {
	return ""
}

func ImplementsType() {
	// 判断是否实现了接口
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false
}

// -----------
func Add(a, b int) int { return a + b }

func invokeFunc() {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 1
}

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
	app := cli.App("defer-demo", "demo for defer usage")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax: 栈关系", cli.ActionCommand(basics))
	app.Command("dr", "defer and return : 谁先被调用呀 : return 之后的语句先执行，defer 后的语句后执行 ", cli.ActionCommand(DeferAndReturn))
	app.Command("ri", "return var scope : return 带命名返回值 会先被初始化为0值 并且整个函数可见", cli.ActionCommand(ReturnValInit))
	app.Command("rd", "return but defer : 带命名返回值 被defer 改写", cli.ActionCommand(NamedReturnVarMeetingDefer))
	app.Command("dp", "defer panic : defer 遇见 panic，但是并不捕获异常的情况", cli.ActionCommand(DeferAndPanic))
	app.Command("dp2", "defer panic : defer 遇见 panic，并捕获异常", cli.ActionCommand(DeferPanicRecover))
	app.Command("dp3", "defer panic : defer 中包含 panic", cli.ActionCommand(DeferIncludePanic))
	app.Command("df", "defer and function : defer 下的函数参数包含子函数", cli.ActionCommand(DeferAndFunc))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	f1 := func() {
		fmt.Println("A")
	}
	f2 := func() {
		fmt.Println("B")
	}
	f3 := func() {
		fmt.Println("C")
	}

	defer f1()
	defer f2()
	defer f3()
}

func DeferAndReturn() {
	deferFunc := func() int {
		fmt.Println("defer func called")
		return 0
	}

	returnFunc := func() int {
		fmt.Println("return func called")
		return 0
	}

	returnAndDefer := func() int {

		defer deferFunc()

		return returnFunc()
	}

	returnAndDefer()
}

func ReturnValInit() {
	// 只要声明函数的返回值变量名称，就会在函数初始化时候为之赋值为 0，而且在函数体作用域可见。
	deferFunc := func(i int) (t int) {

		fmt.Println("t = ", t)

		return 2
	}

	deferFunc(10)
}

func NamedReturnVarMeetingDefer() {
	returnButDefer := func() (t int) { //t初始化0， 并且作用域为该函数全域

		defer func() {
			// 改写返回值   此时覆盖return的那个1 了
			t = t * 10
		}()

		return 1
	}
	fmt.Println(returnButDefer())
}

func DeferAndPanic() {
	defer_call := func() {
		defer func() { fmt.Println("defer: panic 之前1") }()
		defer func() { fmt.Println("defer: panic 之前2") }()

		panic("异常内容") //触发defer出栈

		defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
	}

	defer_call()

	fmt.Println("main 正常结束")

}
func DeferPanicRecover() {
	// defer 最大的功能是 panic 后依然有效 所以 defer 可以保证你的一些资源一定会被关闭，从而避免一些异常出现的问题。
	defer_call := func() {

		defer func() {
			fmt.Println("defer: panic 之前1, 捕获异常")
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

		panic("异常内容") //触发defer出栈

		defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
	}

	defer_call()

	fmt.Println("main 正常结束")
}

func DeferIncludePanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()
	// panic 仅有最后一个可以被 revover 捕获。
	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}

func DeferAndFunc() {
	fn := func(index int, value int) int {

		fmt.Println(index)

		return index
	}
	defer fn(1, fn(3, 0))
	defer fn(2, fn(4, 0))
}

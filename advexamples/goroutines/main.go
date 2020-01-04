package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"runtime"
	"sync"
	"time"
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
	app.Command("bs2", "basic syntax 2", cli.ActionCommand(basicSyntax2))
	app.Command("bs3", "basic syntax 3: 闭包 变量的访问 竞态发生啦", cli.ActionCommand(basicSyntax3))
	app.Command("bs4", "basic syntax 4: 对比上例的不同之处", cli.ActionCommand(basicSyntax4))
	app.Command("bs5", "anonymous func : ", cli.ActionCommand(basicSyntax5))

	app.Command("bp", "best practice : 协程同步的最佳实践 ", cli.ActionCommand(bestPractice))
	app.Command("sd2", "sync demo : 同步例子2  乱序发生了！", cli.ActionCommand(syncDemo2))
	app.Command("at", "available Threads : 机器可用的线程！", cli.ActionCommand(availableThreads))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
func basicSyntax() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}
func basicSyntax2() {
	msg := "hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}
func basicSyntax3() {
	msg := "hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye!"
	time.Sleep(100 * time.Millisecond)
}
func basicSyntax4() {
	msg := "hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
	msg = "Goodbye!"
	time.Sleep(100 * time.Millisecond)
}
func basicSyntax5() {
	msg := "hello"
	go func(msg2 string) {
		fmt.Println(msg2)
	}(msg)
	msg = "Goodbye!"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("hello")
}

var wg = sync.WaitGroup{}
var counter = 0

func bestPractice() {
	var msg = "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye!"
	wg.Wait()
}
func syncDemo2() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello2()
		go increment()
	}
	wg.Wait()
}
func sayHello2() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}
func increment() {
	counter++
	wg.Done()
}

func availableThreads() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

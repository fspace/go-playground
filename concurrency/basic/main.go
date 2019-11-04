package main

import (
	"fmt"
	"playgo/concurrency/basic/lib"
	"time"
)

func main() {
	go fmt.Println("Hello, playground")
	// ----------------------------------------------------------------形式2
	// 方法也是函数哦！
	go lib.A{}.Method()
	// ----------------------------------------------------------------形式3
	//	闭包 Closures are anonymous functions,
	go func() {
		fmt.Println("Hello, playground from closure")
	}()

	time.Sleep(time.Nanosecond)
}

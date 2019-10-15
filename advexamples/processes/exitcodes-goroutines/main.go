package main

import (
	"fmt"
	"os"
	"time"
)

// If the os.Exit function happens in a goroutine, all the goroutines (including the main one) will terminate immediately
// without executing any deferred call, as follows:
func main() {
	go func() {
		defer fmt.Println("go end (deferred)") // 不会被调用
		fmt.Println("go start")
		os.Exit(1)
	}()
	defer fmt.Println("Main end (deferred)!") // 不会被调用
	fmt.Println("Main start")
	time.Sleep(time.Second)
	fmt.Println("main end")
}

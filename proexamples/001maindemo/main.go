package main

import (
	"fmt"
	"os"
)

// @see https://github.com/hashicorp/consul/blob/master/main.go
func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {

	defer func() {
		// 如果捕获错误 那么返回码就是其他的了
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err)
			exitCode = 1
		}
	}()

	fmt.Println("hi  this is fun cli app!")
	panic("hi panic from lalallala ...")

	return
}

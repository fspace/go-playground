package main

import (
	"fmt"
	"playgo/pkg/app/cli"
)

func init() {
	cli.Register("hello-go", helloGo)
}

func main() {
	cli.Init()
	cli.Run()
}

func helloGo() {
	fmt.Println("hello go ")
}

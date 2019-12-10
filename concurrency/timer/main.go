package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(time.Millisecond, func() {
		fmt.Println("Hello 1!")
	})
	t := time.AfterFunc(time.Millisecond*5, func() {
		fmt.Println("Hello 2!")
	})
	if !t.Stop() {
		panic("should not fire")
	}
	time.Sleep(time.Millisecond * 10)
}

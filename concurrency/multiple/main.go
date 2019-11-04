package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 0; i <10; i++ {
	//	go fmt.Println(i)
	//}

	argEval()
	time.Sleep(time.Nanosecond)
}

func argEval() {
	var a int
	// passing value
	go func(v int) { fmt.Println(v) }(a)

	// passing pointer
	go func(v *int) { fmt.Println(*v) }(&a)

	a = 42
}

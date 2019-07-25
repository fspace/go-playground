package main

import (
	"fmt"
	"playgo/examples/package/logic"
)

func main() {
	fmt.Println("_> enter the main thread: ")
	fmt.Println(logic.Add(1, 2))

	var input string
	fmt.Scanln( &input)
	fmt.Println("your input is ",input)
}

package main

import "fmt"

func main() {
	var a = make(chan int)
	a <- 10
	fmt.Println(<-a)
}

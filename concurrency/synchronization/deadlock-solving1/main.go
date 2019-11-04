package main

import "fmt"

func main() {
	var a = make(chan int)
	go func() {
		a <- 10
	}()
	fmt.Println(<-a)
}

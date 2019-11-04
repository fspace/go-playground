package main

import "fmt"

func main() {
	const max = 10
	var a = make(chan int)

	go func() {
		for i := 0; i < max; i++ {
			a <- i
		}
	}()

	for i := 0; i < max; i++ {
		fmt.Println(<-a)
	}
}

package main

import "fmt"

func main() {
	var (
		a = make(chan int, 0)
		b = make(chan int, 5)
	)

	fmt.Println("a is", cap(a))
	fmt.Println("b is", cap(b))

	lenDemo()

	oneWay()
}

func lenDemo() {
	var (
		a = make(chan int, 5)
	)
	for i := 0; i < 5; i++ {
		a <- i
		fmt.Println("a is", len(a), "/", cap(a))
	}
}

func blockingDemo() {
	var (
		a = make(chan int, 5)
	)
	for i := 0; i < 5; i++ {
		a <- i
		fmt.Println("a is", len(a), "/", cap(a))
	}
	a <- 0 // Blocking
}

func oneWay() {
	var a = make(chan int)
	s, r := (chan<- int)(a), (<-chan int)(a)
	fmt.Printf("%T - %T", s, r)
}

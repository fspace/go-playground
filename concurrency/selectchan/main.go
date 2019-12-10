package main

import "time"

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	for i := 0; i < 10; i++ {
		go func() { <-ch1 }()

		go func() { ch2 <- a }()
		time.Sleep(time.Nanosecond) // If we add a very small pause (using time.Sleep) before the select switch,
		// we will have the scheduler pick at least one goroutine and we will then have one of the two operations executed:
		select {
		case ch1 <- b:
			fmt.Println("ch1 got a", b)
		case v := <-ch2:
			fmt.Println("ch2 got a", v)
		default:
			fmt.Println("too slow")
		}
	}
}

func demo2() {
	ch1, ch2 := make(chan int), make(chan int)
	a, b := 2, 10
	go func() { <-ch1 }()
	go func() { ch2 <- a }()
	// A timer exposes a read-only channel, so it's not possible to close it.
	t := time.NewTimer(time.Nanosecond)
	select {
	case ch1 <- b:
		fmt.Println("ch1 got a", b)
	case v := <-ch2:
		fmt.Println("ch2 got a", v)
	case <-t.C:
		fmt.Println("too slow")
	}
}

func timerDemo() {
	t := time.NewTimer(time.Millisecond)
	time.Sleep(time.Millisecond / 2)
	if !t.Stop() {
		panic("it should not fire")
	}
	select {
	case <-t.C:
		panic("not fired")
	default:
		fmt.Println("not fired")
	}
	// -------------------------------------------------------------
	if t.Reset(time.Millisecond) {
		panic("timer should not be active")
	}
	time.Sleep(time.Millisecond)
	if t.Stop() {
		panic("it should fire")
	}
	select {
	case <-t.C:
		fmt.Println("fired")
	default:
		panic("not fired")
	}

}

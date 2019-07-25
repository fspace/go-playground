package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	mu.Lock()

	go func() {
		fmt.Println("hello, world")
		mu.Unlock()
	}()

	mu.Lock()
}

func sync2(){
	done := make(chan int)

	go func() {
		fmt.Println("你好， 世界")
		<-done
	}()

	done <-1
}
// better one
func sync3()  {
	done  := make(chan int, 1)

	go func() {
		fmt.Println("你好，世界")
		done <- 1
	}()

	<-done
}

func printN()  {
	done := make(chan int, 10)

	for i := 0; i< cap(done) ; i++ {
		go func() {
			fmt.Println("你好，世界")
			done <- 1
		}()
	}

	for i := 0; i<cap(done); i++ {
		<-done
	}
}

func wg()  {
	var wg sync.WaitGroup
	for i :=0; i<10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好，世界")
			wg.Done()
		}()
	}

	wg.Wait()
}
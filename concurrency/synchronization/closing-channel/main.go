package main

import "fmt"

func main() {
	const max = 10
	var a = make(chan int)

	go func() {
		for i := 0; i < max; i++ {
			a <- i
		}
		close(a)
	}()
	//for {
	//	v, ok := <-a
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(v)
	//}

	for v := range a {
		fmt.Println(v)
	}
}

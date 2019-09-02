package main

import (
	"fmt"
	"runtime"
	"time"
)

// @see https://mp.weixin.qq.com/s/O3JXr0Zm_c2tkmCtnsaQ6Q?tdsourcetag=s_pcqq_aiomsg

func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	// Set up the pipeline.
	done := make(chan struct{})
	defer close(done)

	out := gen(done, 2, 3)

	for n := range out {
		fmt.Println(n)              // 2
		time.Sleep(5 * time.Second) // done thing, 可能异常中断接收
		if true {                   // if err != nil
			break
		}
	}
}

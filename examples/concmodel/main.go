package main

import "fmt"

// "Do not communicate by sharing memory, share memory by communicating."
// A channel is made for sharing data, and it usually connects two or more execution threads in an application,

// Channel 是用来共享数据的   思考下 原先所有可能被多线程访问的变量 是否都可以替换为一个channel？

func chanExample1() {
	var ch = make(chan int)
	go func() {
		b := <-ch // receive and assign
		fmt.Println(b)
	}()
	ch <- 10 // send to channel
}
func chanExample2() {
	var ch = make(chan int)
	go func() {
		b, ok := <-ch // channel open, ok is true
		if ok {
			fmt.Println("from channel :", b)
		}
		b, ok = <-ch // channel closed, ok is false
		if !ok {
			fmt.Println("channel has been closed!")
		}
		b = <-ch                  // channel close, b will be a zero value
		fmt.Println("b is  :", b) // 但是最好约定不能发0值 不然无法判断此值是因为channel关闭的原因 还是本来对方就发的0值
	}()
	ch <- 10  // send to channel
	close(ch) // close the channel
}

func selectExample() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go func() { ch1 <- 10 }()
	go func() { <-ch2 }()

	var a int
	//switch { // the first operation that completes is selected
	select { // the first operation that completes is selected
	case a = <-ch1:
		fmt.Println("read from ch1:", a)
	case ch2 <- 20:
		fmt.Println("write to ch2:", 20)
	}

}

func main() {
	//chanExample1()
	//chanExample2()

	selectExample()
	for {
	} // 死循环 等待携程结束
}

package main

import "sync"

// 过同步原语来给两个事件明确排序
func main() {
	done := make(chan  int)

	go func() {
		println("你好，世界")
		done <- 1
	}()

	<- done

	sync2()
}

func sync2()  {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		println("你好， 世界2")
		mu.Unlock()
	}()

	mu.Lock()
}
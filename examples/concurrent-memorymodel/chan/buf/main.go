package main

// 可以根据控制Channel的缓存大小来控制并发执行的Goroutine的最大数目,
var limit = make(chan int, 3)

func main() {
	work := []func(){}
	work = append(work, func() {
        println("work 1")
	})
	work = append(work, func() {
		println("work 2")
	})
	work = append(work, func() {
		println("work 3")
	})
	for _, w := range work{
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}

}

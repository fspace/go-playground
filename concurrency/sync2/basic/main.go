package basic

func main() {
	a := make(chan int)
	go send(a, 10)
	done := make(chan struct{})
	go receive(a, done)
	<-done
}

// ----------------------------------------------------
func send(ch chan<- int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println(v)
	}
	close(done)
}

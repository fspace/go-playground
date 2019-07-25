package main

var done = make(chan bool)
var msg string

func aGoruotine()  {
	msg = "你好，世界"
	// 若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值。
	close(done)
	//done <- true
}


func main() {
	go aGoruotine()
	<- done
	println(msg)
}

package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("Starting the application")

	c := make(chan os.Signal)
	signal.Notify(c) // 接收所以信号
	s := <-c
	log.Println("Exit with signal :", s)
}

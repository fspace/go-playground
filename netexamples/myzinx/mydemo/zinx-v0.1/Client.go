package main

import (
	"log"
	"net"
	"time"
)

// main 模拟客户端
func main() {
	log.Println("client start...")

	time.Sleep(time.Second * 1)
	//
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Fatalln("Client start err ", err, " exit!")
	}
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1..."))
		if err != nil {
			log.Fatalln("write conn err :", err)
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Fatalln("read buf err: ", err)
		}
		log.Printf("server call back : %s, cnt= %d \n", buf, cnt)

		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}

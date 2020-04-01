package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"playgo/netexamples/myzinx/internal/znet"
	"time"
)

// main 模拟客户端
func main() {
	log.Println("client0 start...")

	time.Sleep(time.Second * 1)
	//
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Fatalln("Client start err ", err, " exit!")
	}
	for {
		// 发送封包的Msg MsgID:0
		dp := znet.NewDataPack()
		binMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.7 	Client Test Message	")))
		if err != nil {
			fmt.Println("Pack error :", err)
			return
		}
		if _, err = conn.Write(binMsg); err != nil {
			fmt.Println("write error:", err)
			return
		}

		// ## 服务器就应该给我们回复一个msg 数据
		// 先读取流中的head部分 得到ID和dataLend
		binHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binHead); err != nil {
			fmt.Println("read head err:", err)
			break
		}
		// 将二进制的head拆包至msg结构体中
		msgHead, err := dp.Unpack(binHead)
		if err != nil {
			fmt.Println("client unpack msgHead error: ", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			// msg里有数据啦
			// 再跟进dataLen 进行第二次读取 将data读出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error,", err)
				return
			}
			fmt.Println("---> Recv Server Msg: ID=", msg.Id,
				", len", msg.DataLen,
				", data=", string(msg.Data))
		}

		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}

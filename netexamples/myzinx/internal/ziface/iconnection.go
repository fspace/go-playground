package ziface

import "net"

type IConnection interface {
	// 启动 让当前连接准备开始工作
	Start()
	// 停止 结束当前连接的工作
	Stop()
	// 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接的连接ID
	GetConnID() uint32
	// 获取远程客户端 TCP的状态  IP:PORT
	RemoteAddr() net.Addr
	// 发送数据， 将数据发送给远程的客户端
	// Send(data []byte) error
	SendMsg(msgId uint32, data []byte) error
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error

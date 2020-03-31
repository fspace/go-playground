package znet

import (
	"fmt"
	"log"
	"net"
	"playgo/netexamples/myzinx/internal/ziface"
)

type Connection struct {
	Conn   *net.TCPConn
	ConnID uint32

	isClosed bool

	// 当前连接所绑定的处理业务的方法API
	handleAPI ziface.HandleFunc

	// 告知当前连接已经退出的/停止 channel
	ExitChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn,
	connID uint32,
	callback_api ziface.HandleFunc,
) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	log.Println("Reader Goroutine is running ...")
	defer fmt.Println("connID= ", c.ConnID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buffer中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("new buf err:", err)
			continue
		}

		// 调用当前连接所绑定的HandleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID ", c.ConnID, "handle is err: ", err)
			break
		}

	}
}
func (c *Connection) Start() {
	log.Println("Connection Start()... ConnID=", c.ConnID)
	// 启动从当前连接的读数据的业务
	go c.StartReader()
	// TODO 启动从当前连接写数据的业务
}

func (c *Connection) Stop() {
	log.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true
	// 关闭socket
	c.Conn.Close()
	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (*Connection) Send(data []byte) error {
	panic("implement me")
}

var _ ziface.IConnection = &Connection{}

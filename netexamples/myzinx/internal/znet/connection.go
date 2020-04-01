package znet

import (
	"errors"
	"fmt"
	"io"
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

	// 该连接处理的方法router
	Router ziface.IRouter
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn,
	connID uint32,
	//callback_api ziface.HandleFunc,
	router ziface.IRouter,
) *Connection {
	c := &Connection{
		Conn:   conn,
		ConnID: connID,
		//handleAPI: callback_api,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	log.Println("Reader Goroutine is running ...")
	defer fmt.Println("connID= ", c.ConnID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buffer中
		//buf := make([]byte, 512)
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		////cnt, err := c.Conn.Read(buf)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("new buf err:", err)
		//	continue
		//}
		// 创建拆包解包对象
		dp := NewDataPack()
		// 读取客户端的MSG Head 二进制流的 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("Read msg head err: ", err)
			break
		}
		// 拆包 得到MsgID 放在消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack err:", err)
			break
		}
		// 根据dataLen 再次读取Data 放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error:", err)
				break
			}
		}
		msg.SetData(data)

		req := Request{
			conn: c,
			//data: buf,
			msg: msg,
		}
		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		//// 调用当前连接所绑定的HandleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnID ", c.ConnID, "handle is err: ", err)
		//	break
		//}

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

//func (*Connection) Send(data []byte) error {
//	panic("implement me")
//}

// 提供一个SendMsg方法 将我们要发送给客户端的数据先进行封包 再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg \n")
	}
	// 讲data进行封包 MsgDataLen MsgID MsgData
	dp := NewDataPack()
	binMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error messg id = ", msgId)
		return errors.New("Pack error msg")
	}
	// 讲数据发送给客户端
	if _, err := c.Conn.Write(binMsg); err != nil {
		fmt.Println("Write msg id ", msgId, " error ", err)
		return errors.New("conn Write error")
	}
	return nil
}

var _ ziface.IConnection = &Connection{}

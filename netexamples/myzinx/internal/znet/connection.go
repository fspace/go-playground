package znet

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"playgo/netexamples/myzinx/internal/ziface"
	"playgo/netexamples/myzinx/utils"
	"sync"
)

type Connection struct {
	// 当前连接归属的Server
	TcpServer ziface.IServer

	Conn   *net.TCPConn
	ConnID uint32

	isClosed bool

	// 当前连接所绑定的处理业务的方法API
	// handleAPI ziface.HandleFunc

	// 告知当前连接已经退出的/停止 channel (reader告知writer 退出  reader是处理的上游)
	ExitChan chan bool

	// 无缓冲的管道， 用于读,写Goroutine之间的消息通讯
	msgChan chan []byte

	// 该连接处理的方法router
	// Router ziface.IRouter

	// 消息的管理MsgId合对应的处理业务API 关系
	MsgHandle ziface.IMsgHandle

	// 连接属性集合
	property map[string]interface{}
	// 保护连接属性的锁
	propertyLock sync.RWMutex
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, errors.New("no property found ")
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	// 删除属性
	delete(c.property, key)
}

// 初始化连接模块的方法
func NewConnection(
	server ziface.IServer,
	conn *net.TCPConn,
	connID uint32,
	//callback_api ziface.HandleFunc,
	//router ziface.IRouter,
	msgHandler ziface.IMsgHandle,
) *Connection {
	c := &Connection{
		TcpServer: server,
		Conn:      conn,
		ConnID:    connID,
		//handleAPI: callback_api,
		// Router:   router,
		isClosed: false,
		msgChan:  make(chan []byte),
		ExitChan: make(chan bool, 1),

		MsgHandle: msgHandler,

		property: make(map[string]interface{}),
	}
	// 将连接加入连接管理器
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

func (c *Connection) StartReader() {
	log.Println("[Reader Goroutine is running ...]")
	defer fmt.Println("[Reader is exit], connID= ", c.ConnID, " remote addr is ", c.RemoteAddr().String())
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
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandle.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandle.DoMsgHandler(&req)
		}

		// 执行注册的路由方法
		//go func(request ziface.IRequest) {
		//	//c.Router.PreHandle(request)
		//	//c.Router.Handle(request)
		//	//c.Router.PostHandle(request)
		//
		//}(&req)
		//go c.MsgHandle.DoMsgHandler(&req)

		//// 调用当前连接所绑定的HandleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnID ", c.ConnID, "handle is err: ", err)
		//	break
		//}

	}
}

// 写消息的Goroutine 专门发送给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println("[Conn Writer exit!]", c.RemoteAddr().String())
	// 不断的阻塞 等待channel的消息 进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			// 有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error:", err)
				return
			}
		case <-c.ExitChan:
			// 代表Reader已经退出 此时Writer 也要退出
			return
		}
	}
}

func (c *Connection) Start() {
	log.Println("Connection Start()... ConnID=", c.ConnID)
	// 启动从当前连接的读数据的业务
	go c.StartReader()
	// 启动从当前连接写数据的业务
	go c.StartWriter()

	// 执行钩子方法
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	log.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	// 销毁连接之前执行的业务Hook
	c.TcpServer.CallOnConnStop(c)

	// 关闭socket
	c.Conn.Close()

	// 告知Writer关闭
	c.ExitChan <- true

	// 将当前连接从连接管理器中移除
	c.TcpServer.GetConnMgr().Remove(c)

	// 回收资源
	close(c.ExitChan)
	close(c.msgChan)
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
	fmt.Println("Connection::SendMsg: ID=", msgId)

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
	//if _, err := c.Conn.Write(binMsg); err != nil {
	//	fmt.Println("Write msg id ", msgId, " error ", err)
	//	return errors.New("conn Write error")
	//}
	c.msgChan <- binMsg

	return nil
}

var _ ziface.IConnection = &Connection{}

package znet

import (
	"errors"
	"fmt"
	"log"
	"net"
	"playgo/netexamples/myzinx/internal/ziface"
	"playgo/netexamples/myzinx/utils"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int

	// 给当前server添加router server注册的连接对应的处理业务
	// Router ziface.IRouter
	// 当前server的消息管理模块， 用来邦定MsgID 对应的业务处理逻辑
	MsgHandler ziface.IMsgHandle
	// 该server 的连接管理器
	ConnMgr ziface.IConnManager
	// ## 事件支持
	OnConnStart func(conn ziface.IConnection)
	OnConnStop  func(conn ziface.IConnection)
}

// -------------------------------------------------------------------------
// ## 事件支持
// SetOnConnStart 设置连接开始的回调方法
func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("----> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> Call OnConnStop() ...")
		s.OnConnStop(conn)
	}
}

// -------------------------------------------------------------------------

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Succ!!!")
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s, Listenner at IP: %s, Port: %d is Starting \n ",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version %s, MaxConn: %d , MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize,
	)

	log.Println("Start me")
	go func() {
		// 开启消息队列以及Worker工作池
		s.MsgHandler.StartWorkerPool()
		start(s)
	}()
}

// 定义当前客户端连接所绑定的handle_api（目前写死）
func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallbackClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err :", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func start(s *Server) {
	log.Println("Start me")
	// 1. 获取一个TCP的地址
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Fatal("resolve tcp addr error: ", err)
	}
	// 2. 监听服务器地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		log.Fatalln("listen: ", s.IPVersion, "err ", err)
	}
	log.Println("start zinx server succ.", s.Name, "succ. Listening ...")

	var cid uint32
	cid = 0
	// 3.阻塞等待客户端连接 ， 处理客户端连接业务(读写）
	for {
		conn, err := listenner.AcceptTCP()
		if err != nil {
			log.Println("Accept err: ", err)
			continue
		}
		// 连接限制检查
		if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
			// TODO 给客户端响应 连接过多的错误包
			fmt.Println("====== > Too many Connection ,MaxConn = ", utils.GlobalObject.MaxConn)
			conn.Close()
			continue
		}
		// 将处理心连接的业务方法 和 Conn进行绑定 得到我们的连接模块
		//dealConn := NewConnection(conn, cid, CallbackToClient)
		dealConn := NewConnection(s, conn, cid, s.MsgHandler)
		cid++
		go dealConn.Start()
		// 客户端已经连接 做一些业务 做一个最基本的最大512字节长度的回显业务
		//go func() {
		//	for {
		//		buf := make([]byte, 512)
		//		cnt, err := conn.Read(buf)
		//		if err != nil {
		//			log.Println("recv buf err: ", err)
		//			continue
		//		}
		//		log.Printf("recv client buf %s, cnt %d\n", buf, cnt)
		//		// 回显功能
		//		if _, err := conn.Write(buf[:cnt]); err != nil {
		//			log.Println("Write back buf err :", err)
		//			continue
		//		}
		//	}
		//}()

	}
}

func (s *Server) Stop() {
	fmt.Println("[Stop] Zinx server name :", s.Name)
	// panic("implement me")
	//  将一些服务器的资源，状态或者已经开辟的连接信息 进行停止或者回收
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}
func NewServer(name string) ziface.IServer {
	// 构造器惯用法 还有类似的 NewXxxWithYYY
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		//Router:    nil,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

// 可以渐进式创建server 逐个给成员赋值 不同的WithXxx
func (s *Server) WithPort(port int) *Server {
	s.Port = port
	return s
}

// ===================================================================
// 事件的方法支持

// ===================================================================

var _ ziface.IServer = &Server{}

package znet

import (
	"errors"
	"fmt"
	"log"
	"net"
	"playgo/netexamples/myzinx/internal/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	log.Println("Start me")
	go func() {
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
		// 将处理心连接的业务方法 和 Conn进行绑定 得到我们的连接模块
		dealConn := NewConnection(conn, cid, CallbackToClient)
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

func (*Server) Stop() {
	// panic("implement me")
	// TODO 将一些服务器的资源，状态或者已经开辟的连接信息 进行停止或者回收
}

func (s *Server) Serve() {
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

func NewServer(name string) ziface.IServer {
	// 构造器惯用法 还有类似的 NewXxxWithYYY
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}

// 可以渐进式创建server 逐个给成员赋值 不同的WithXxx
func (s *Server) WithPort(port int) *Server {
	s.Port = port
	return s
}

var _ ziface.IServer = &Server{}

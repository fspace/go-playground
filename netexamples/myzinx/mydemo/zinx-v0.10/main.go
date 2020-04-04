package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"playgo/netexamples/myzinx/internal/ziface"
	"playgo/netexamples/myzinx/internal/znet"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("run", "run the server", cli.ActionCommand(Run))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
// 自定义路由 ping逻辑处理
type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handler...")
	fmt.Println("recv from client: MsgID=  ", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping... ping... ping...ddddd"))
	if err != nil {
		fmt.Println(err)
	}
}

// ..................
type HelloZinxRouter struct {
	znet.BaseRouter
}

func (pr *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handler...")
	fmt.Println("recv from client: MsgID=  ", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome to Znix!!"))
	if err != nil {
		fmt.Println(err)
	}
}

// ==============================================================================
// 创建连接之后执行的钩子函数
func DoConnBegin(conn ziface.IConnection) {
	fmt.Println("===> DoConnBegin is Called!")
	if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}
	// 给当前连接设置一些属性
	fmt.Println("Set conn properties: eg. Name , Home...")
	conn.SetProperty("Name", "丹丹--冰冰")
	conn.SetProperty("Home", "https://github.com/aceId")
	conn.SetProperty("SomeKey", "some value")
}

// 连接断开之前需要执行的函数
func DoConnStop(conn ziface.IConnection) {
	fmt.Println("===> DoConnStop is Called...")
	fmt.Println("conn ID = ", conn.GetConnID(), "is lost ...")

	// 获取连接属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}
	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Home = ", home)
	}
	if someKey, err := conn.GetProperty("SomeKey"); err == nil {
		fmt.Println("SomeKey = ", someKey)
	}

}

// ==============================================================================
func Run() {
	// 1. 创建server
	s := znet.NewServer("[zinx V0.10]")

	// 注册连接相关的钩子
	s.SetOnConnStart(DoConnBegin)
	s.SetOnConnStop(DoConnStop)

	// 注册处理路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	// 3. 启动server
	s.Serve()
}

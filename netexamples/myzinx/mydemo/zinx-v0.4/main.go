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

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandler...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("call back before ping err :", err)
	}
}
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handler...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping..."))
	if err != nil {
		fmt.Println("call back  ping err :", err)
	}
}
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandler...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..."))
	if err != nil {
		fmt.Println("call back after ping err :", err)
	}
}

// ==============================================================================
func Run() {
	// 1. 创建server
	s := znet.NewServer("[zinx V0.4]")

	// 注册处理路由
	s.AddRouter(&PingRouter{})
	// 3. 启动server
	s.Serve()
}

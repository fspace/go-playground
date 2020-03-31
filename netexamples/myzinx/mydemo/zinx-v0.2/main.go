package main

import (
	"github.com/jawher/mow.cli"
	"os"
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

// ==============================================================================
func Run() {
	// 1. 创建server
	s := znet.NewServer("[zinx v0.2]")
	s.Serve()
}

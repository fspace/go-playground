package main

import (
	"github.com/jawher/mow.cli"
	"os"
	"playgo/netexamples/hinet/internal/basic_1"
	"playgo/netexamples/hinet/internal/basic_2"
	"playgo/netexamples/hinet/internal/basic_3"
	"playgo/netexamples/hinet/internal/concurrency_client"
	"playgo/netexamples/hinet/internal/concurrency_server"
	"playgo/netexamples/hinet/internal/udp_client"
	"playgo/netexamples/hinet/internal/udp_client2"
	"playgo/netexamples/hinet/internal/udp_server"
	"playgo/netexamples/hinet/internal/udp_server2"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("go-net-demo", "demo for net programing")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic tcp server", cli.ActionCommand(basic_1.Main))
	app.Command("bs2", "basic tcp server : 带响应功能的简单服务器", cli.ActionCommand(basic_3.Main))
	app.Command("bc", "basic tcp client", cli.ActionCommand(basic_2.Main))

	app.Command("cs", "concurrency server", cli.ActionCommand(concurrency_server.Main))
	app.Command("cc", "concurrency client: 可以在终端输入要发送的数据", cli.ActionCommand(concurrency_client.Main))

	app.Command("us", "udp server", cli.ActionCommand(udp_server.Main))
	app.Command("uc", "udp client", cli.ActionCommand(udp_client.Main))

	app.Command("us2", "udp server v2", cli.ActionCommand(udp_server2.Main))
	app.Command("uc2", "udp client v2: 可以在终端输入要发送的数据哦", cli.ActionCommand(udp_client2.Main))
	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {

}

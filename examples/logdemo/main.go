package main

import (
	"github.com/jawher/mow.cli"
	"os"
	"playgo/examples/logdemo/internal/hilogrus"
	"playgo/examples/logdemo/internal/hilogrus/customlogger"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("log-demo", "demo for log")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))

	app.Command("logrus:basic", "logrus 基本用法", cli.ActionCommand(hilogrus.Main))
	app.Command("logrus:custom", "logrus 定制化", cli.ActionCommand(customlogger.Main))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {

}

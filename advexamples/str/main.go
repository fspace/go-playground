package main

import (
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("string-demo", "demo for string related functions ")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {

}

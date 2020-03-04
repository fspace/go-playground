package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"playgo/webexamples/gininaction/internal/milestones/basic"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("gin-in-action", "build web app with gin framework")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	// app.Command("bs", "basic syntax", cli.ActionCommand(basics))

	app.Command("bs", "basic syntax", cli.ActionCommand(basic.Main))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	fmt.Println("hi I am a skeleton function !")
}

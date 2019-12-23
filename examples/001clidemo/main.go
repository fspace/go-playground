package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("uman", "User Manager")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman get"
	app.Command("get", "get a user details", func(cmd *cli.Cmd) {
		var (
			detailed = cmd.BoolOpt("detailed", false, "Disaply detailed info")
			id       = cmd.StringArg("ID", "", "The user id to display")
		)

		cmd.Action = func() {
			fmt.Printf("user %q details (detailed mode: %v)\n", *id, *detailed)
		}
	})
	// Declare command, which is invocable with "uman info"
	app.Command("info", "show some information", cli.ActionCommand(info))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)
}

// =================================================================================================
// ## helper func
func info() {
	fmt.Println("function info  invoked! ")
	panic("panic from func info! ")
}

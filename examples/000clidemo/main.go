package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("uman", "User Manager")

	// ---------------------------------------------------------------------------------------
	//			## 可选设置

	app.Spec = "[-v]"

	var (
		verbose = app.BoolOpt("v verbose", false, "Verbose debug mode")
	)

	app.Before = func() {
		if *verbose {
			// Here we can enable debug output in our logger for example
			fmt.Println("Verbose mode enabled")
		}
	}
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare our first command, which is invocable with "uman list"
	app.Command("list", "list the users", func(cmd *cli.Cmd) {
		// These are the command-specific options and args, nicely scoped
		// inside a func so they don't pollute the namespace
		var (
			all = cmd.BoolOpt("all", false, "List all users, including disabled")
		)

		// Run this function when the command is invoked
		cmd.Action = func() {
			// Inside the action, and only inside, we can safely access the
			// values of the options and arguments
			fmt.Printf("user list (including disabled ones: %v)\n", *all)
		}
	})

	// Declare our second command, which is invocable with "uman get"
	app.Command("get", "get a user details", func(cmd *cli.Cmd) {
		var (
			detailed = cmd.BoolOpt("detailed", false, "Disaply detailed info")
			id       = cmd.StringArg("ID", "", "The user id to display")
		)

		cmd.Action = func() {
			fmt.Printf("user %q details (detailed mode: %v)\n", *id, *detailed)
		}
	})

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)
}

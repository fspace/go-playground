package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

// Global options available to any of the commands
var filename *string

// TODO 未处理参数传递哦  自己去官网看看 cmd.XxxOpt  cmd.XxxArg 系列方法
func main() {
	app := cli.App("crud-demo", "this is a crud demo")

	// Define our top-level global options
	filename = app.StringOpt("f file", os.Getenv("HOME")+"/.safe", "Path to safe")

	// Define our command structure for usage like this:
	//app.Command("list", "list accounts", cmdList)
	//app.Command("creds", "display account credentials", cmdCreds)
	app.Command("user", "manage user", func(config *cli.Cmd) {
		uc := &userCmd{}
		config.Command("list", "list users", uc.List)
		config.Command("add", "add an user", uc.Add)
		config.Command("update", "update an user", uc.Update)
		config.Command("get", "get an user by id", uc.Get)
		config.Command("remove", "remove an user(s)", uc.Remove)
	})

	app.Run(os.Args)
}

// =========================================================================

type userCmd struct {
}

func (c *userCmd) List(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("list users")
	}
}

func (c *userCmd) Add(cmd *cli.Cmd) {
	// cmd.Spec = "ACCOUNT [ -u=<username> ] [ -p=<password> ]"
	var (
		username = cmd.StringOpt("u username", "admin", "Account username")
		password = cmd.StringOpt("p password", "admin", "Account password")
	)
	cmd.Action = func() {
		fmt.Printf("add users")
		fmt.Println("username: ", *username)
		fmt.Println("password: ", *password)
	}
}
func (c *userCmd) Update(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("update user")
	}
}
func (c *userCmd) Get(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("get user")
	}
}
func (c *userCmd) Remove(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("remove users")
	}
}

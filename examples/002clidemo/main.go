package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

/**
本例演示 如何用类似注册表的风格来注册函数  所有的注册全部用注册方法 在init中来做  也可以单独提取为一个公共包 做为全局package
*/
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

	// 加载命令表中的entry
	if len(actionCommands) > 0 {
		for k, v := range actionCommands {
			app.Command(k, v.Desc, cli.ActionCommand(v.Fn))
		}
	}

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)
}

// ------------------------------------------------------------------------

type actionCmdInfo struct {
	Desc string
	Fn   func()
}

// 注册表
var actionCommands map[string]actionCmdInfo

func registerActionCMDFunc(k string, cmdInfo actionCmdInfo) {
	// 签名参数也可以改为 k string desc string, cmdFn func() 三个参数 这样actionCmdInfo类型就不必让使用者知道了！

	if actionCommands == nil {
		actionCommands = make(map[string]actionCmdInfo)
	}

	if _, exists := actionCommands[k]; exists {
		panic(fmt.Errorf("func %q is already registered", k))
	}
	actionCommands[k] = cmdInfo
}

func init() {
	registerActionCMDFunc("info2", actionCmdInfo{Desc: "print some information", Fn: info2})
}

// =================================================================================================
// ## helper func
func info() {
	fmt.Println("function info  invoked! ")
	// panic("panic from func info! ")
}

func info2() {
	fmt.Println("function info2  invoked! ")
}

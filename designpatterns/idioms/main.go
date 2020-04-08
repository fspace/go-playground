package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"playgo/designpatterns/idioms/internal/options"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("go idioms demo", "demo for idioms usage")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "构造方法 灵活初始化属性", cli.ActionCommand(basics))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	u := options.NewUser(options.WithName("Yiqing"))

	fmt.Printf("user: %#v\n", u)
	u2 := options.NewUser2(options.WithName("Yiqing2"))
	fmt.Printf("user: %#v\n", u2)

	u3 := options.NewUser3(options.WithSex(1), options.WithName("yiqing3"))
	fmt.Printf("user: %#v\n", u3)
}

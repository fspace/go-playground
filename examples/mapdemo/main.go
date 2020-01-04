package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	//basic2()
	//ptrArray()
	realmain()
}
func realmain() {
	app := cli.App("map-demo", "demo for map type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("demo1", "map 变量的引用复制", cli.ActionCommand(assignment))


	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)
}

// ==============================================================================
func assignment()  {
	a := map[string]string{"foo": "bar", "baz":"buz"}
	b := a // 引用复制
	fmt.Println(a , b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

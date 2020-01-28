package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"playgo/books/levelup/internal/ch1"
	"playgo/books/levelup/internal/ch2"
	"playgo/books/levelup/internal/ch3"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))

	app.Command("ch1/tc", "Numeric Types Conversion: 数字类型的转换", cli.ActionCommand(ch1.NumericTypesConversion))
	app.Command("ch1/as-len", "Arrays And Slices : len : 内建len函数", cli.ActionCommand(ch1.ArraysAndSlices2))
	app.Command("ch1/as-loop", "Arrays And Slices : looping : iterate over a slice,", cli.ActionCommand(ch1.Looping))
	app.Command("ch1/map", "Map : ", cli.ActionCommand(ch1.Maps))
	app.Command("ch1/map2", "Map2 : ", cli.ActionCommand(ch1.Map2))
	app.Command("ch1/df", "Define Func : ", cli.ActionCommand(ch1.DefineFunc))
	app.Command("ch1/sc", "StructsCreating : ", cli.ActionCommand(ch1.StructsCreating))
	app.Command("ch1/tm", "Type Methods : ", cli.ActionCommand(ch1.TypeMethods))
	app.Command("ch1/tm2", "Type Methods2 : 类型方法 接受者可以是引用类型", cli.ActionCommand(ch1.TypeMethods2))

	app.Command("ch2/ct", "Custom Types : 自定义类型", cli.ActionCommand(ch2.CustomTypes))
	app.Command("ch2/ii", "InterfacesImplements : 接口的隐式实现", cli.ActionCommand(ch2.InterfacesImplements))
	app.Command("ch2/err", "DoSomeWebRequest : 做一些web请求 处理http返回码 ", cli.ActionCommand(ch2.DoSomeWebRequest))
	app.Command("ch2/et", "EmbeddedTypes : 类型内嵌 ", cli.ActionCommand(ch2.EmbeddedTypes))
	app.Command("ch2/defer", "CopyFile : defer 示例 ", cli.ActionCommand(func() {
		err := ch2.CopyFile("dst.go.me", "main.go")
		if err != nil {
			fmt.Println(err)
		}
	}))
	app.Command("ch2/md", "Blackfriday :  ", cli.ActionCommand(ch2.Blackfriday))

	app.Command("ch3/ss", "simple web server :  ", cli.ActionCommand(ch3.SimpleServer))
	// curl localhost:3000 -i
	app.Command("ch3/ss2", "simple web server : 操纵Header部分  ", cli.ActionCommand(ch3.SimpleServer2))
	app.Command("ch3/ps", "PathAndSubtrees :   ", cli.ActionCommand(ch3.PathAndSubtrees))
	app.Command("ch3/ps", "HttpStatusHelpers :   ", cli.ActionCommand(ch3.HttpStatusHelpers))
	app.Command("ch3/hd", "HandlerDemo :   ", cli.ActionCommand(ch3.HandlerDemo))
	app.Command("ch3/mw", "MiddlewareDemo :   ", cli.ActionCommand(ch3.MiddlewareDemo))
	app.Command("ch3/ht", "HtmlTemplates :   ", cli.ActionCommand(ch3.HtmlTemplates))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {

}

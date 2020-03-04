package main

import (
	"encoding/json"
	"fmt"
	"github.com/jawher/mow.cli"
	"io/ioutil"
	"os"
	"playgo/books/levelup/internal/ch1"
	"playgo/books/levelup/internal/ch2"
	"playgo/books/levelup/internal/ch3"
	"playgo/books/levelup/internal/ch4"
	"playgo/books/levelup/internal/ch5"
	"playgo/books/levelup/internal/ch6"
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
	app.Command("ch3/tc", "Html TemplateConditionals :   ", cli.ActionCommand(ch3.TemplateConditionals))
	app.Command("ch3/lr", "template  LoopsWithRange :   ", cli.ActionCommand(ch3.TemplateLoopsWithRange))
	app.Command("ch3/par", "template  reuse :   ", cli.ActionCommand(ch3.TemplatePartials))
	app.Command("ch3/pip", "template  pipeline :   ", cli.ActionCommand(ch3.TemplatePipelines2))
	app.Command("ch3/tv", "template  variable :   ", cli.ActionCommand(ch3.TemplateVariables))
	app.Command("ch3/jm", "JsonMarshal :   ", cli.ActionCommand(ch3.JsonMarshal))
	app.Command("ch3/cjk", "JsonMarshal : CustomJSONKeys  ", cli.ActionCommand(ch3.CustomJSONKeys))
	app.Command("ch3/nt", "JsonMarshal : NestedTypes  ", cli.ActionCommand(ch3.NestedTypes))

	app.Command("ch3/ut", "Unmarshaling Types  ", cli.ActionCommand(func() {
		conf := ch3.Config{}
		// 对于文件跟结构体失配的地方 额外key 会被忽略
		data, err := ioutil.ReadFile("config/app.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(data, &conf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Site: %s (%s)", conf.Name, conf.URL)

		db := conf.Database
		// Print out a database connection string.
		fmt.Printf(
			"DB: mysql://%s:%s@%s:%d/%s",
			db.Username,
			db.Password, db.Host,
			db.Port,
			db.Name,
		)
	}))
	app.Command("ch3/ujs", "Unmarshaling: UnknownJSONStructure   ", cli.ActionCommand(ch3.UnknownJSONStructure))

	app.Command("ch4/sw", " Simple static webserver  ", cli.ActionCommand(ch4.SimpleStaticWebserver))
	app.Command("ch4/ss", " SimpleServer  ", cli.ActionCommand(ch4.SimpleServer))
	app.Command("ch5/s", " SimpleServer  ", cli.ActionCommand(ch5.Main))
	app.Command("ch6/s", " SimpleServer  ", cli.ActionCommand(ch6.Main))
	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {

}

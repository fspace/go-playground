package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"playgo/webexamples/gininaction/internal/actions/grouproute"
	"playgo/webexamples/gininaction/internal/actions/jsonoutput"
	"playgo/webexamples/gininaction/internal/actions/queryarray"
	"playgo/webexamples/gininaction/internal/actions/queryform"
	"playgo/webexamples/gininaction/internal/actions/renderhtml"
	"playgo/webexamples/gininaction/internal/actions/routeparams"
	"playgo/webexamples/gininaction/internal/actions/simplerestful"
	"playgo/webexamples/gininaction/internal/actions/urlparams"
	"playgo/webexamples/gininaction/internal/actions/xmloutput"
	"playgo/webexamples/gininaction/internal/milestones/basic"
	"playgo/webexamples/gininaction/internal/milestones/hellogin"
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

	app.Command("bs", "basic web app", cli.ActionCommand(basic.Main))
	app.Command("gin:hello", "basic gin app ", cli.ActionCommand(hellogin.Main))
	app.Command("gin:restful", "simple gin restful app ", cli.ActionCommand(simplerestful.Main))
	app.Command("gin:route-params", "simple gin restful app : 可变的路由参数 ", cli.ActionCommand(routeparams.Main))
	// /users/123/go      匹配
	app.Command("gin:route-params2", "simple gin restful app : 可变的路由参数 : 任意匹配 ", cli.ActionCommand(routeparams.MainAsterisk))
	app.Command("gin:query-params", "gin query string : 获取查询参数 http://localhost:8080/?wechat=some_key", cli.ActionCommand(urlparams.Main))
	app.Command("gin:query-params2", "gin query string : 获取查询参数 如果不传递key则给出默认值 http://localhost:8080/?wechat=some_key", cli.ActionCommand(urlparams.MainDefaultQueryParams))
	app.Command("gin:query-array", "gin query string : 接收数组和 Map http://localhost:8080/?media=blog&media=wechat",
		cli.ActionCommand(queryarray.Main))
	app.Command("gin:query-map ", "gin query string : 接收数组和 Map http://localhost:8080/map?ids[a]=123&ids[b]=456&ids[c]=789",
		cli.ActionCommand(queryarray.MainQueryMap))
	app.Command("gin:query-form ", "gin query handle : 处理表单 |  访问：curl -d wechat=some-key http://localhost:8080/ ",
		cli.ActionCommand(queryform.Main))

	app.Command("gin:group-route ", "gin route : 分组路由 |  访问：  http://localhost:8080/v1/users ",
		cli.ActionCommand(grouproute.Main))
	app.Command("gin:group-route2 ", "gin route : 分组路由中间件 |  访问：  http://localhost:8080/v2/users ",
		cli.ActionCommand(grouproute.Main_GroupRoute_Midleware))
	app.Command("gin:group-route3 ", "gin route : 分组路由嵌套 |  访问：  http://localhost:8080/v1/admin/users ",
		cli.ActionCommand(grouproute.Main_GroupRoute_Nested))

	app.Command("gin:json-output ",
		"gin render : json输出 |  访问：  http://localhost:8080/hello ",
		cli.ActionCommand(jsonoutput.Main))
	app.Command("gin:json-output2 ",
		"gin render : 结构体转json |  访问：  http://localhost:8080/users/123 ",
		cli.ActionCommand(jsonoutput.Main_Struct2JSON))
	app.Command("gin:template ",
		"gin tpl : html 模板输出 |  访问：  http://localhost:8080 ",
		cli.ActionCommand(renderhtml.Main_stdlib))
	app.Command("gin:template2 ",
		"gin tpl : gin 中模板的使用 |  访问：  http://localhost:8080/html ",
		cli.ActionCommand(renderhtml.Main_ginHtml))
	app.Command("gin:template3 ",
		"gin tpl : gin 中模板的使用 : 加载目录 |  访问：  http://localhost:8080/html ",
		cli.ActionCommand(renderhtml.Main_loadFiles))

	app.Command("gin:template4 ",
		"gin tpl : gin 中模板的使用 : 自定义函数 |  访问：  http://localhost:8080/html ",
		cli.ActionCommand(renderhtml.Main_customFuncs))

	app.Command("gin:render-xml ",
		"gin xml output : gin 中渲染xml输出  |  访问：  http://localhost:8080/xml ",
		cli.ActionCommand(xmloutput.Main))

	app.Command("gin:render-xml2 ",
		"gin xml output : gin 中渲染xml输出 自定义结构体 |  访问：  http://localhost:8080/xml ",
		cli.ActionCommand(xmloutput.Main_customStruct))

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

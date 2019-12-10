package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/imdario/mergo" // 相当好用的库 被多个著名项目使用哦  用法如： https://github.com/godep-migrator/rigger-host/blob/master/host/config.go
	"playgo/webexamples/web-basic/lib/templatedemo"
	"strconv"
	"strings"

	//"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并解析模板内容
		//tmpl, err := template.New("test").Parse("Hello world!")
		tplStr := `Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}

SiteName : {{ .SiteName }}
CurrentUser: {{.User}}
`
		// 模板也可以逐步构造："golang 几种字符串的连接方式" https://segmentfault.com/a/1190000012978989
		tplStr = `
{{$name := "Alice"}}
{{$age := 18}}
{{$round2 := true}}
Name: {{$name}}
Age: {{$age}}
Round2: {{$round2}}

{{$name = "Bob"}}
Name: {{$name}}
`
		tmpl, err := template.New("test").Parse(tplStr)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 数据准备
		type BaseEnv struct {
			SiteName string // 网站名称
			User     string // 当前用户  在真实场景中可以是一个对象哦
		}
		be := BaseEnv{
			SiteName: "my-site",
			User:     "yiqing", // 可以来自session哦
		}
		baseData := struct2map(be)

		// 调用模板对象的渲染方法
		//err = tmpl.Execute(w, nil)
		currentData := map[string]interface{}{
			"SKU":       "xxx",
			"Name":      "hello",
			"UnitPrice": 1,
			"Quantity":  5,
		}
		// 跟基础数据做合并
		mergo.Merge(&baseData, currentData) // 和jquery  extend  有点像呢！

		fmt.Printf("%#v", baseData)
		err = tmpl.Execute(w, baseData)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	// localhost:4000/div?x=1&y=2
	http.HandleFunc("/div", divHandler)
	http.HandleFunc("/compare", compareHandler)
	http.HandleFunc("/range", rangeDemo)
	http.HandleFunc("/map", mapDemo)
	http.HandleFunc("/with", withDemo)
	http.HandleFunc("/scope", withScope)
	http.HandleFunc("/func", tplFuncDemo)
	http.HandleFunc("/pipeline", tplFuncPipelineDemo)
	http.HandleFunc("/pipeline2", tplFuncPipelineDemo2)
	http.HandleFunc("/tpl-reuse", tplReuse)
	http.HandleFunc("/local-tpl", localTpl())
	http.HandleFunc("/home", localTplHome())
	http.HandleFunc("/text-tpl", templatedemo.TextTemplateDemo)

	http.HandleFunc("/xss", templatedemo.XssDemo)
	http.HandleFunc("/xss2", templatedemo.XssDemo2)

	http.HandleFunc("/safe-html", templatedemo.SafeHtmlDemo)
	http.HandleFunc("/bluemonday", templatedemo.BluemondayDemo())
	http.HandleFunc("/deli", templatedemo.DelimsDemo)

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

// ------------------------------------------------------------------------------------ handlers
func divHandler(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
{{if .yIsZero}}
	除数不能为 0
{{else}}
	{{.result}}
{{end}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 获取 URL 查询参数的值
	// 注意：为了简化代码逻辑，这里并没有进行错误处理
	x, _ := strconv.ParseInt(r.URL.Query().Get("x"), 10, 64)
	y, _ := strconv.ParseInt(r.URL.Query().Get("y"), 10, 64)

	// 当 y 不为 0 时进行除法运算
	yIsZero := y == 0
	result := 0.0
	if !yIsZero {
		result = float64(x) / float64(y)
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"yIsZero": yIsZero,
		"result":  result,
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
{{$name1 := "alice"}}
{{$name2 := "bob"}}
{{$age1 := 18}}
{{$age2 := 23}}

{{if eq $age1 $age2}}
	年龄相同
{{else}}
	年龄不相同
{{end}}

{{if ne $name1 $name2}}
	名字不相同
{{end}}

{{if gt $age1 $age2}}
	alice 年龄比较大
{{else}}
	bob 年龄比较大
{{end}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

func rangeDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
{{range $name := .Names}}
	{{$name}}
{{end}}

{{range $i, $name := .Names}}
	{{$i}}. {{$name}}
{{end}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"Names": []string{
			"Alice",
			"Bob",
			"Carol",
			"David",
		},
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

func mapDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
{{range $name, $val := .}}
	{{$name}}: {{$val}}
{{end}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"Names": []string{
			"Alice",
			"Bob",
			"Carol",
			"David",
		},
		"Numbers": []int{1, 3, 5, 7},
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func withDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`Inventory
SKU: {{.Inventory.SKU}}
Name: {{.Inventory.Name}}
UnitPrice: {{.Inventory.UnitPrice}}
Quantity: {{.Inventory.Quantity}}

========================

{{with .Inventory}}
	SKU: {{.SKU}}
	Name: {{.Name}}
	UnitPrice: {{.UnitPrice}}
	Quantity: {{.Quantity}}
{{end}}

================================
blank
{{- with .Inventory}}
	SKU: {{.SKU}}
	Name: {{.Name}}
	UnitPrice: {{.UnitPrice}}
	Quantity: {{.Quantity}}
{{- end}}

`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"Inventory": Inventory{
			SKU:       "11000",
			Name:      "Phone",
			UnitPrice: 699.99,
			Quantity:  666,
		},
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

func withScope(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
{{$name1 := "alice"}}
name1: {{$name1}}
{{with true}}
	{{$name1 = "alice2"}}
	{{$name2 := "bob"}}
	name2: {{$name2}}
{{end}}
name1 after with: {{$name1}}

`)

	// 注意同上例的细微区别 :=   =
	tmpl, err = template.New("test").Parse(`
{{$name1 := "alice"}}
name1: {{$name1}}
{{with true}}
	{{$name1 := "alice2"}}
	{{$name2 := "bob"}}
	name1 in with: {{$name1}}
	name2: {{$name2}}
{{end}}
name1 after with: {{$name1}}
`)

	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

func tplFuncDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并添加自定义模板函数
	tmpl := template.New("test").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})

	// 解析模板内容
	_, err := tmpl.Parse(`
result: {{add 1 2}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}
func tplFuncPipelineDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并添加自定义模板函数
	tmpl := template.New("test").Funcs(template.FuncMap{
		"add2": func(a int) int {
			return a + 2
		},
	})

	// 解析模板内容
	_, err := tmpl.Parse(`
result: {{add2 0 | add2 | add2}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}
func tplFuncPipelineDemo2(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并添加自定义模板函数
	tmpl := template.New("test").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})

	// 解析模板内容
	_, err := tmpl.Parse(`
result: {{add 1 3 | add 2 | add 2}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}
func tplReuse(w http.ResponseWriter, r *http.Request) {
	/**
		   - 通过 Funcs 方法添加了名为 join 模板函数，其实际上就是调用 strings.Join
	  	   - 通过 define "<名称>" 的语法定义了一个非常简单的局部模板，即以根对象 . 作为参数调用 join 模板函数
		   - 通过 template "<名称>" <参数> 的语法，调用名为 list 的局部模板，并将 .names 作为参数传递进去（传递的参数会成为局部模板的根对象）
	*/

	// 创建模板对象并添加自定义模板函数
	tmpl := template.New("test").Funcs(template.FuncMap{
		"join": strings.Join,
	})

	// 解析模板内容
	_, err := tmpl.Parse(`
{{define "list"}}
    {{join . ", "}}
{{end}}
Names: {{template "list" .names}}
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"names": []string{"Alice", "Bob", "Cindy", "David"},
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

// 本地模板文件
func localTplHome() http.HandlerFunc {
	// The name of the template is the bare filename of the template, not the complete path Execute will execute the
	// default template provided it’s named to match.
	// 模板默认名称是第一个参数的文件名（非文件全路径 只是取裸名称 具体细节可以看源码) 如果多个模板文件同名那么采用后者 前面的被“遮盖"

	// 创建模板对象并解析模板内容
	tmpl, err := template.ParseFiles("views/_base.tmpl", "views/home.html") // 接受多个模板文件哦 模板文件间的关系可以是 Master-nest之类

	// template.ParseGlob(pattern)会根据pattern解析所有匹配的模板并保存。

	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	log.Println("模板文件名称：", tmpl.Name())
	return func(w http.ResponseWriter, r *http.Request) {

		// 渲染指定模板的内容
		err = tmpl.ExecuteTemplate(w, "_base.tmpl", map[string]interface{}{
			"names": []string{"Alice", "Bob", "Cindy", "David"},
		})

		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	}
}
func localTpl() http.HandlerFunc {
	// 创建模板对象并解析模板内容
	tmpl, err := template.ParseFiles("views/template_local.tmpl") // 接受多个模板文件哦 模板文件间的关系可以是 Master-nest之类
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, map[string]interface{}{
			"names": []string{"Alice", "Bob", "Cindy", "David"},
		})
		//      // 渲染指定模板的内容
		//		err = tmpl.ExecuteTemplate(w, "template_local.tmpl", map[string]interface{}{
		//			"names": []string{"Alice", "Bob", "Cindy", "David"},
		//		})

		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	}
}
func localTpl2() http.HandlerFunc {

	/**
		注意：

	   - 函数的注入，必须要在parseFiles之前，因为解析模板的时候，需要先把函数编译注入。
	   - Template object can have multiple templates in it and each one has a name. If you look at the implementation of
		ParseFiles, you see that it uses the filename as the template name inside of the template object.
		So, name your file the same as the template object, (probably not generally practical) or else use ExecuteTemplate instead of just Execute.
	   - The name of the template is the bare filename of the template, not the complete path。如果模板名字写错了，
		执行的时候会出现：”**” is an incomplete or empty template
		————————————————————————————————————————————————————————————————
		原文链接：https://blog.csdn.net/wj199395/article/details/75040723
	*/

	// 注册模板函数
	// 创建模板对象并添加自定义模板函数
	tmpl := template.New("template_local.tmpl").Funcs(template.FuncMap{
		"join": strings.Join,
	})
	// 创建模板对象并解析模板内容
	//tmpl, err := template.ParseFiles("views/template_local.tmpl")
	tmpl, err := tmpl.ParseFiles("views/template_local.tmpl")

	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, map[string]interface{}{
			"names": []string{"Alice", "Bob", "Cindy", "David"},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	}
}

// ------------------------------------------------------------------------------------ handlers.

// 结构体转Map
func struct2map(s interface{}) map[string]interface{} {
	// https://stackoverflow.com/questions/23589564/function-for-converting-a-struct-to-map-in-golang
	return structs.Map(s)
	// 这个方法也挺好的： FillMap 可能也可以完成map合并功能
}

package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/imdario/mergo" // 相当好用的库 被多个著名项目使用哦
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
		mergo.Merge(&baseData, currentData)

		fmt.Printf("%#v", baseData)
		err = tmpl.Execute(w, baseData)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

// 结构体转Map
func struct2map(s interface{}) map[string]interface{} {
	// https://stackoverflow.com/questions/23589564/function-for-converting-a-struct-to-map-in-golang
	return structs.Map(s)
}

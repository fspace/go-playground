package templatedemo

import (
	"fmt"
	"html/template"
	"net/http"
)

// 修改模板定界符
func DelimsDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Delims("[[", "]]").Parse(`[[.content]]`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"content": "Hello world!",
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

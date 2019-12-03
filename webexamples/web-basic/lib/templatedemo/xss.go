package templatedemo

import (
	"fmt"
	h_template "html/template" // 此包会产生 跨站脚本攻击
	"net/http"
	"text/template" // 此包会产生 跨站脚本攻击
)

/**
按照官方的说法，html/template 本身是一个 text/template 包的一层封装，并在此基础上专注于提供安全保障。作为使用者来说，最直观的变化就是对所有的文本变量都进行了转义处理
*/

func XssDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
<html>
<body>
	<p>{{.content}}</p>
</boyd>
</html>
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"content": "<script>alert('you have been pwned')</script>",
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}
func XssDemo2(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := h_template.New("test").Parse(`
<html>
<body>
	<p>{{.content}}</p>
</boyd>
</html>
`)
	if err != nil {
		fmt.Fprintf(w, "Parse: %v", err)
		return
	}

	// 调用模板对象的渲染方法
	err = tmpl.Execute(w, map[string]interface{}{
		"content": "<script>alert('you have been pwned')</script>",
	})
	if err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

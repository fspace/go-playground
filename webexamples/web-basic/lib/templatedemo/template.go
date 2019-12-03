package templatedemo

import (
	"fmt"
	"net/http"
	"text/template"
)

func TextTemplateDemo(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并解析模板内容
	tmpl, err := template.New("test").Parse(`
<html>
<body>
	<h2>Heading 2</h2>
	<p>Paragraph</p>
</boyd>
</html>
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

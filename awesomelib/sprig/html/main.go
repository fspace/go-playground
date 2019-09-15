package main

import (
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"path/filepath"
)

func main() {

	pattern := filepath.Join( /*".",*/ "tpls", "*.html") // 找某个目录下所有模板文件
	// This example illustrates that the FuncMap *must* be set before the
	// templates themselves are loaded.
	tpl := template.Must(
		// name很诡异 要跟文件名一致么？ 不然报错呀
		template.New("tpl1.html").Funcs(sprig.FuncMap()).ParseGlob(pattern))

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		// 下面的错误源于文件名 模板名称跟文件名不对应？
		// template: "base" is an incomplete or empty template
		panic(err)
	}
}

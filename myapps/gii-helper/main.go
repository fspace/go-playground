package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/go-playground/validator.v9"
)

// bool 类型的选项
var (
	h bool
)

// options
var (
	db    string
	table string
)

//// flag 太多时 对于不同用途的选项 也可以归组选项 不同用途的归为一个类型 这样变量就有了group了 便于管理和区分
// 对了 验证库对结构好像有更好的亲和 虽然也支持验证变量 :)   结构的话支持tag验证规则的配置
//type option struct {
//	db    string
//	table string
//	// toleration time.Duration
//}
//var opt &option{}
// 	opt := &option{}
//	flag.StringVar(&opt.db, "db", "", "set db name")

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	// 另一种绑定方式
	// q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	// flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")

	flag.StringVar(&db, "d", "", "set db name`")
	flag.StringVar(&table, "t", "", "set table name")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

/**
@see https://www.jianshu.com/p/f9cf46a4de0e
@see https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter13/13.1.html

命令行 flag 的语法有如下三种形式：

-flag // 只支持bool类型
-flag=x
-flag x // 只支持非bool类型
*/
func main() {
	flag.Parse()
	if h {
		flag.Usage()
	}

	//// ## 验证flags
	//db2 := ""
	//errs := validate.Var(&db2, "required")
	//if errs != nil {
	//	fmt.Println(errs) //
	//	return
	//}
	////db ok, move on
	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `gii-helper version: gii-helper/0.0.1
Usage: gii-helper [-h] [-d dbName] [-t tableName] 

Options:
`)
	flag.PrintDefaults()
}

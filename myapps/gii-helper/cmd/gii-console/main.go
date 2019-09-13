package main

import (
	"flag"
	"fmt"
	"playgo/myapps/gii-helper/pkg"
	"playgo/myapps/gii-helper/pkg/util"

	//"log"
	log "github.com/sirupsen/logrus" // replace the std log package
	"os"
	// "gopkg.in/go-playground/validator.v9" // 竟然用不了！

	"github.com/go-ozzo/ozzo-validation"
)

// TODO 下版本 把所有的flag 变量整理到struct 这样利于配置验证规则 一个个验证太麻烦了
// TODO 也可以考虑暴露为 web-server程序 便于和其他程序交互 现在是命令行程序
// bool 类型的选项
var (
	h bool
)

// options
var (
	db    string // db name
	table string // table name
)

/**
   // FIXME 常规配置不能满足要求 那么久需要特定配置  只需要配置driver 跟 DataSource 就可以了 但需要参考手册自己拼字符串 属于高级用法
   // TODO 有空了 改为DBOptions  不然变量太零散 看着难受
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "your_password"
	dbName="your_db_name"
*/
var (
	dbDriver     = "mysql"
	dbDataSource string // "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8"

	// FIXME 下面跟上面的dataSource 是冗余配置 底下是简易配置 上面是高级配置 应对特殊情况
	dbUser = "root"
	dbPass = ""
	dbHost = "127.0.0.1"
	dbPort = 3306
)

// logging
var (
	debug bool
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

	// ## 设置日志级别
	flag.BoolVar(&debug, "debug", false, "enable the log output")

	// 数据库连接
	flag.StringVar(&dbDriver, "dbDriver", "mysql", "set the database driver: mysql|sqlite3|postgres refer to xorm driver supporting")
	flag.StringVar(&dbDataSource, "ds", "", "set the dataSourceName . refer to go sql.Open method ")

	flag.StringVar(&dbUser, "du", dbUser, "set the db user . eg root ")
	flag.StringVar(&dbPass, "ps", dbPass, "set the db password . ")
	flag.StringVar(&dbHost, "dh", dbHost, "set the db host . eg: localhost 127.0.0.1  ")
	flag.IntVar(&dbPort, "dp", dbPort, "set the db port .")

}

// use a single instance of Validate, it caches struct info
//var validate *validator.Validate

/**
@see https://www.jianshu.com/p/f9cf46a4de0e
@see https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter13/13.1.html

命令行 flag 的语法有如下三种形式：

-flag // 只支持bool类型
-flag=x
-flag x // 只支持非bool类型

lugrus 用法： @see http://xiaorui.cc/2018/01/11/golang-logrus%E7%9A%84%E9%AB%98%E7%BA%A7%E9%85%8D%E7%BD%AEhook-logrotate/

*/
func main() {
	// usage-sample： go run main.go -d acontent -t ac_config --debug

	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	if !debug {
		log.SetLevel(log.WarnLevel)
	}

	//// ## 验证flags
	//db2 := ""
	//errs := validate.Var(&db2, "required")
	//if errs != nil {
	//	fmt.Println(errs) //
	//	return
	//}
	////db ok, move on
	err := validation.Errors{
		"dbName":    validation.Validate(db, validation.Required),
		"tableName": validation.Validate(table, validation.Required),
		//"zip": validation.Validate(c.Address.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	}.Filter()
	if err != nil {
		fmt.Println("\r\n ", err, "\r\n ")
		usage()
		return
	}
	// validation ok, move on !

	var ds string
	if dbDataSource != "" {
		ds = dbDataSource
	} else {
		if dbDriver == "mysql" {
			// ds = fmt.Sprintf("root:@/%s?charset=utf8", db)
			// "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8"
			ds = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, db)
		} else if dbDriver == "postgres" {
			ds = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				dbHost, dbPort, dbUser, dbPass, db)
		} else {
			log.Fatalln("you should specify the <DataSource> flag !")
		}

	}
	//
	itr := pkg.NewDBInteractor(pkg.DBOption{
		DriverName: dbDriver,
		// "root:@/test?charset=utf8"
		DSName: ds,
	})
	rslt := itr.GetColumnsForTable(table)
	// PrettyPrint(rslt)
	util.PrintJson(rslt)

}

func usage() {
	fmt.Fprintf(os.Stderr, `gii-console version: gii-console/0.0.1
Usage: gii-console [-h] [-d dbName] [-t tableName] 

Options:
`)
	flag.PrintDefaults()
}

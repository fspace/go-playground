package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	//"log"
	log "github.com/sirupsen/logrus" // replace the std log package
	"os"
	"xorm.io/core"

	// "gopkg.in/go-playground/validator.v9" // 竟然用不了！

	"github.com/go-ozzo/ozzo-validation"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// TODO 下版本 把所有的flag 变量整理到struct 这样利于配置验证规则 一个个验证太麻烦了
// TODO 也可以考虑暴露为 web-server程序 便于和其他程序交互 现在是命令行程序
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
//var validate *validator.Validate

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

	//
	itr := NewDBInteractor(DBOption{
		DriverName: "mysql",
		// "root:@/test?charset=utf8"
		DSName: fmt.Sprintf("root:@/%s?charset=utf8", db), // FIXME 这里可能需要暴露给flag 需要数据库的用户名跟密码 现在写死了！
	})
	rslt := itr.GetColumnsForTable(table)
	// PrettyPrint(rslt)
	PrintJson(rslt)

}

func usage() {
	fmt.Fprintf(os.Stderr, `gii-helper version: gii-helper/0.0.1
Usage: gii-helper [-h] [-d dbName] [-t tableName] 

Options:
`)
	flag.PrintDefaults()
}

// =========================================================================  +|
// ##              core engin        -------------  +|
//             TODO 有空了提出到其他目录去
//

func NewDBInteractor(opt DBOption) *DBInteractor {
	inst := &DBInteractor{}
	inst.Option = opt
	return inst
}

type DBOption struct {
	DriverName string // DriverName: mysql
	DSName     string // DataSourceName
}

type DBInteractor struct {
	Option    DBOption
	XormEngin xorm.Engine
}

type MyColumn struct {
	core.Column
	GoType string
}

func (itr *DBInteractor) GetColumnsForTable(name string) map[string]*MyColumn /**core.Column*/ {
	var err error
	//	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	//engine, err := xorm.NewEngine(Config.GetString("db.driver", "mysql"),
	//	Config.GetString("db.dataSourceName", "root:@/test?charset=utf8"))
	engine, err := xorm.NewEngine(itr.Option.DriverName,
		itr.Option.DSName)
	checkErr(err)
	err = engine.Ping()
	checkErr(err)

	/*
		db := engine.DB()
		tables := db.
	*/
	dlc := engine.Dialect()
	log.Println(" db name : ", dlc.URI().DbName)
	tables, err := dlc.GetTables()
	checkErr(err)

	var tbl *core.Table
	for _, t := range tables {

		// for i, t := range tables {
		/*
			log.Printf("\n <--    table:%d    \t name: %s    --> \n", i, tbl.Name)
			colSeq, cols, err := dlc.GetColumns(tbl.Name)
			checkErr(err)
			PrettyPrint(colSeq)
			for nm, col := range cols {
				// PrettyPrint(col)
				fmt.Printf("\n\n name: %s  \t sql-type: %s  \t go-type: %s \n",
					nm,
					col.SQLType.Name,
					core.SQLType2Type(col.SQLType).Name())
			}
		*/
		if t.Name == name {
			tbl = t
			break
		}
	}
	if tbl == nil {
		log.Println("no such table :", name)
		// panic(name + " does not exists !")
		return nil // TODO 后期需要返回特定结构啦！
	}

	// 处理列
	colSeq, cols, err := dlc.GetColumns(tbl.Name)
	checkErr(err)

	var _ = colSeq
	//	PrettyPrint(colSeq)

	fmt.Printf("\n  name  \t  sql-type  \t  go-type  \n")
	fmt.Printf("================================================")

	var results = make(map[string]*MyColumn, len(cols))

	for nm, col := range cols {
		// PrettyPrint(col)
		fmt.Printf("\n  %s  \t  %s  \t  %s  ",
			nm,
			col.SQLType.Name,
			core.SQLType2Type(col.SQLType).Name())

		results[nm] = &MyColumn{
			Column: *col,
			GoType: core.SQLType2Type(col.SQLType).Name(),
		}

	}
	fmt.Println("\n")

	return results //cols

}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func PrettyPrint(v interface{}) {
	//   fmt.Printf("%#v", p) //with name, value and type
	// b, _ := json.MarshalIndent(v, "", "  ")
	// println(string(b))
	spew.Dump(v)
}
func PrintJson(v interface{}) {
	//   fmt.Printf("%#v", p) //with name, value and type
	b, _ := json.MarshalIndent(v, "", "  ")
	//fmt.Println(b)
	fmt.Println(string(b))
}

package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus" // replace the std log package
	"os"
	"playgo/myapps/gii-helper/pkg/util"
	"xorm.io/core"
)

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
	Option DBOption
	//XormEngin xorm.Engine
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
	util.CheckErr(err)

	// ## 设置xorm日志
	f, err := os.Create("sql.log")
	util.CheckErr(err)
	engine.SetLogger(xorm.NewSimpleLogger(f))

	err = engine.Ping()
	util.CheckErr(err)

	/*
		db := engine.DB()
		tables := db.
	*/
	dlc := engine.Dialect()
	log.Println(" db name : ", dlc.URI().DbName)
	tables, err := dlc.GetTables()
	util.CheckErr(err)

	var tbl *core.Table
	for _, t := range tables {

		// for i, t := range tables {
		/*
			log.Printf("\n <--    table:%d    \t name: %s    --> \n", i, tbl.Name)
			colSeq, cols, err := dlc.GetColumns(tbl.Name)
			util.CheckErr(err)
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
	util.CheckErr(err)

	var _ = colSeq
	//	PrettyPrint(colSeq)

	log.Printf("\n  name  \t  sql-type  \t  go-type  \n")
	log.Printf("================================================")

	var results = make(map[string]*MyColumn, len(cols))

	for nm, col := range cols {
		// PrettyPrint(col)
		log.Printf("\n  %s  \t  %s  \t  %s  ",
			nm,
			col.SQLType.Name,
			core.SQLType2Type(col.SQLType).Name())

		results[nm] = &MyColumn{
			Column: *col,
			GoType: core.SQLType2Type(col.SQLType).Name(),
		}

	}
	log.Println("\n")

	return results //cols

}

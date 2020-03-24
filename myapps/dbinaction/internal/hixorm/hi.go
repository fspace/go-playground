package hixorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"

	"xorm.io/core"
	"xorm.io/xorm"
)

type User struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Birthday time.Time `xorm:"DATE"`
	Sex      string    `xorm:"CHAR(1)"`
	Address  string    `xorm:"VARCHAR(256)"`
}

func Main() {

	//创建orm引擎
	engine, err := xorm.NewEngine("mysql", "root:@/playgo?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return
	}
	//连接测试
	if err := engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})

	//--------------------------------------------------------------------
	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}
	//--------------------------------------------------------------------

	//增
	//user := new(User)
	//user.Username="tyming"
	//affected,err := engine.Insert(user)
	//fmt.Println(affected)

	//删
	//user := new(User)
	//user.Username="tyming"
	//affected_delete,err := engine.Delete(user)
	//fmt.Println(affected_delete)

	//改
	//user := new(User)
	//user.Username="tyming"
	//affected_update,err := engine.Id(1).Update(user)
	//fmt.Println(affected_update)

	//查
	user := new(User)
	//result,err := engine.Id(1).Get(user)
	result, err := engine.Where("id=?", 1).Get(user)
	fmt.Println(result)

}

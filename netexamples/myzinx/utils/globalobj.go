package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"playgo/netexamples/myzinx/internal/ziface"
)

type GlobalObj struct {
	/**
	Server
	*/
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	/**
	zinx
	*/
	Version        string
	MaxConn        int    // 当前服务器运行的最大连接数
	MaxPackageSize uint32 // 当前zinx框架数据包的最大值
}

// 提供一个全局对外的对象
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		log.Fatal("loadFile err: ", err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func init() {
	// 如果配置文件没有加载  默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerAPP",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}

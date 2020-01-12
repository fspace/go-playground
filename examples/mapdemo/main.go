package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/kr/pretty"
	"os"
	"playgo/examples/mapdemo/internal/services"
)

// Maps are Go’s representation of hash tables, a data structure that allows you to map one arbitrary data type to another.

func main() {
	//basic2()
	//ptrArray()
	realmain()
}
func realmain() {
	app := cli.App("map-demo", "demo for map type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("demo1", "map 变量的引用复制", cli.ActionCommand(assignment))

	app.Command("bs", "map基础知识", cli.ActionCommand(basics))
	app.Command("m2i", "MappingStrings2Interfaces: 字符串到接口的映射", cli.ActionCommand(MappingStrings2Interfaces))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)
}

// ==============================================================================
func assignment() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a // 引用复制
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

func creation() {
	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}

	fmt.Println(youtubeSubscribers["MKBHD"])
}

func basics() {
	// a map of string to int which has
	// no set capacity
	mymap := make(map[string]int)

	// a map of bool to int which has a
	// set capacity of 2
	//boolmap := make(map[bool]int)

	mymap["mykey"] = 10
	fmt.Println(mymap["mykey"]) // prints out 10

	//## Iterating over Keys and Values
	for key, value := range mymap {
		fmt.Println(key)
		fmt.Println(value)
	}

	// 获取所有的keys
	var keyArray []string

	for key := range mymap {
		keyArray = append(keyArray, key)
	}
	pretty.Print(keyArray)

	// ## 删除
	delete(mymap, "mykey")
	fmt.Println("Value deleted from map")
	pretty.Println(mymap)
}

func MappingStrings2Interfaces() {
	fmt.Println("Go Maps Tutorial")
	// we can define a map of string uuids to
	// the interface type 'Service'
	interfaceMap := make(map[string]services.Service)

	// we can then populate our map with
	// simple ids to particular services
	interfaceMap["SERVICE-ID-1"] = services.MyService{}
	interfaceMap["SERVICE-ID-2"] = services.SecondService{}

	// Incoming HTTP Request wants service 2
	// we can use the incoming uuid to lookup the required
	// service and call it's SayHi() method
	interfaceMap["SERVICE-ID-2"].SayHi()

	for key, service := range interfaceMap {
		fmt.Println(key)
		service.SayHi()
	}
}

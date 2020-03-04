package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
)

/**
- golang 避坑指南(1)interface 之坑多多
*/
type Talkable interface {
	TalkEnglish(string)
	TalkChinese(string)
}

type Student1 struct {
	Talkable // 内嵌接口  按需实现
	Name     string
	Age      int
}

func (s *Student1) TalkEnglish(s1 string) {
	fmt.Printf("I'm %s,%d years old,%s", s.Name, s.Age, s1)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic : ", r)
		}
	}()
	//a := Student1{Name: "aaa", Age: 12}
	//a.TalkEnglish("nice to meet you\n")
	//
	//a.TalkChinese("汉语呀 你会说么")

	os.Exit(realMain())
}

func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	// app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("al", "alias", cli.ActionCommand(EmptyInterfaceAlias))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ========================================= =======================================================
type object interface{}
type baseObj = interface{}

func EmptyInterfaceAlias() {
	Foo1("yes i am a string")
	Foo2(2)
}
func Foo1(obj object) {
	fmt.Printf("Type: %T => Value: %v \n", obj, obj)
}
func Foo2(obj baseObj) {
	fmt.Printf("Type: %T => Value: %v \n", obj, obj)
}

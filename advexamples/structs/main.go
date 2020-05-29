package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"os"
	"reflect"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("s", "print the struct info ", cli.ActionCommand(actionStruct))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Age2 ageint `json:"age"`
}

// ==============================================================================
func basics() {
	s := Student{}
	rt := reflect.TypeOf(s)
	f, ok := rt.FieldByName("Name")
	if !ok {
		fmt.Println("no field Name")
		return
	}

	fmt.Println(f.Tag.Get("json"))
}

type ageint int

func (ageint) Tag() string {
	return "age"
}

type taggable interface {
	Tag() string
}

//var _ taggable = &ageint()

func actionStruct() {
	obj := Student{}
	// Create a new struct type:
	s := structs.New(obj)
	//n := structs.Names(server)
	names := s.Names()
	for _, n := range names {
		fmt.Println(n)
		f := s.Field(n)
		// 遍历每个filed
		// Check if the field is exported or not
		if f.IsExported() {
			fmt.Println(n, " field is exported")
		}

		// 查看当前field的值 是否实现了某个接口
		if tagger, ok := f.Value().(taggable); ok {
			fmt.Println(n, "tage is : ", tagger.Tag())
		}

	}
}

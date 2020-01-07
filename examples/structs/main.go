package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/kr/pretty" // 漂亮打印 go结构体 调试用
	"os"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("struct", "demo for studying struct ")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("ns", "Nested Structs", cli.ActionCommand(NestedStructs))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
// our Person struct
type Person struct {
	name string
	age  int
}

func basics() {
	// declaring a new `Person`
	//var myPerson Person

	// declaring a new `elliot`
	elliot := Person{name: "Elliot", age: 24}

	// trying to roll back time to before I was injury prone
	elliot.age = 18

	fmt.Printf("%T: %v \n", elliot, elliot)
}

func NestedStructs() {
	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)

	fmt.Printf("%# v", pretty.Formatter(celtic))
}

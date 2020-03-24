package main

import (
	"fmt"
	// "flag"
	"github.com/jawher/mow.cli"
	"github.com/jinzhu/configor"
	"os"
	"playgo/myapps/dbinaction/cmd/db-inaction/config"
	"playgo/myapps/dbinaction/internal/hellodb"
	"playgo/myapps/dbinaction/internal/hixorm"
)

func main() {
	os.Exit(realMain())
}

// Version indicates the current version of the application.
var Version = "1.0.0"

// var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func realMain() (exitCode int) {
	// 加载配置
	cfg := &config.Config{}
	configor.Load(&cfg, "../../config/app.yml")
	fmt.Printf("config: %#v \n\n", cfg)

	app := cli.App("db-demo", "demo for  the db in go ")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("hello-db", "Accessing the Database", cli.ActionCommand(basics))
	app.Command("query", "Fetching Data from the Database", cli.ActionCommand(hellodb.Action_FetchingData))
	app.Command("query2", "Preparing Queries: You should, in general, always prepare queries to be used multiple times", cli.ActionCommand(hellodb.Action_PreparingQueries))
	app.Command("query3", "Single-Row Queries: a query returns at most one row", cli.ActionCommand(hellodb.Action_SingleRowQueries))
	app.Command("exec:insert", "Statements that Modify Data: insert", cli.ActionCommand(hellodb.Action_StatementsThatModifyData))
	app.Command("error", "Errors From QueryRow()", cli.ActionCommand(hellodb.Action_ErrorsFromQueryRow))
	app.Command("unknown-cols", "Working with Unknown Columns", cli.ActionCommand(hellodb.Way2))

	// ## xorm
	app.Command("xorm:hi", "Xorm : sync2db ", cli.ActionCommand(hixorm.Main))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	hellodb.Main()
}

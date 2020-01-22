package main

import (
	"bytes"
	"github.com/jawher/mow.cli"
	"os"
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

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
// ## ACTION CMD

func basics() {

}

// ==============================================================================
// ## lib  or  util|helpers function
func template(text string, data map[string]interface{}) string {
	// Get a new instance of the template engine
	t := template.New("template")

	// parse the template text:
	tree, _ := t.Parse(text)

	// Execute the template engine on the provided data
	var out bytes.Buffer
	tree.Execute(&out, data)

	return out.String()
}

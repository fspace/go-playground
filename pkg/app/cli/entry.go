package cli

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
)

var globalCli *cli.CLI

func Init() {
	globalCli = cli.NewCLI("app", "1.0.0")
	globalCli.Args = os.Args[1:]
}

func Run() {
	globalCli.Commands = Map()

	exitStatus, err := globalCli.Run()
	if err != nil {
		log.Println(err)
	}

	// fmt.Println("exit-status: ", exitStatus)
	os.Exit(exitStatus)
}

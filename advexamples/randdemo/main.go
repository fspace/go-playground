package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"math/rand"
	"os"
	"time"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("random-number-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("r2", "random in range", cli.ActionCommand(rangeRand))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {

	nanotime := int64(time.Now().Nanosecond())
	rand.Seed(nanotime)

	//
	for i := 0; i < 10; i++ {
		fmt.Printf("random float: %2.2f\n ", 100*rand.Float32())
		fmt.Printf("random int: %d \n ", rand.Int())
	}

}

func rangeRand() {
	nanotime := int64(time.Now().Nanosecond())
	rand.Seed(nanotime)
	for i := 0; i < 500; i++ {
		fmt.Println("rand int[0---100):", rand.Intn(100))
	}

}

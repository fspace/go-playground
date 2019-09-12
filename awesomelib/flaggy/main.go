package main

import "github.com/integrii/flaggy"

var stringFlag = "defaultValue"

func init() {
	subcommand := flaggy.NewSubcommand("subcommandExample")
	subcommand.String(&stringFlag, "f", "flag", "A test string flag")
	flaggy.AttachSubcommand(subcommand, 1)
	flaggy.Parse()
}

// go  run main.go subcommandExample -f test
func main() {
	print(stringFlag)
}

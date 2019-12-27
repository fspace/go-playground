package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// @see https://github.com/etcd-io/bbolt/blob/master/cmd/bbolt/main.go
var (
	// ErrUsage is returned when a usage message was printed and the process
	// should simply exit with an error.
	ErrUsage = errors.New("usage")

	// ErrUnknownCommand is returned when a CLI command is not specified.
	ErrUnknownCommand = errors.New("unknown command")

	// ... 各种 error！
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err == ErrUsage {
		os.Exit(2)
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Main represents the main program execution.
type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// NewMain returns a new instance of Main connect to the standard input/output.
func NewMain() *Main {
	return &Main{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// Run executes the program.
func (m *Main) Run(args ...string) error {
	// Require a command at the beginning.
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		fmt.Fprintln(m.Stderr, m.Usage())
		return ErrUsage
	}

	// Execute command.
	switch args[0] {
	case "help":
		fmt.Fprintln(m.Stderr, m.Usage())
		return ErrUsage

	case "info":
		return newInfoCommand(m).Run(args[1:]...)
	//case "stats":
	//	return newStatsCommand(m).Run(args[1:]...)
	default:
		return ErrUnknownCommand
	}
}

// Usage returns the help message.
func (m *Main) Usage() string {
	return strings.TrimLeft(`
Bolt is a tool for inspecting bolt databases.
Usage:
	bolt command [arguments]
The commands are:
    bench       run synthetic benchmark against bolt
    buckets     print a list of buckets
    check       verifies integrity of bolt database
    compact     copies a bolt database, compacting it in the process
    dump        print a hexadecimal dump of a single page
    get         print the value of a key in a bucket
    info        print basic info
    keys        print a list of keys in a bucket
    help        print this screen
    page        print one or more pages in human readable format
    pages       print list of pages with their types
    page-item   print the key and value of a page item.
    stats       iterate over all pages and generate usage stats
Use "bolt [command] -h" for more information about a command.
`, "\n")
}

// =====================================================================================
// ## COMMANDS

// InfoCommand represents the "info" command execution.
type InfoCommand struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// NewInfoCommand returns a InfoCommand.
func newInfoCommand(m *Main) *InfoCommand {
	return &InfoCommand{
		Stdin:  m.Stdin,
		Stdout: m.Stdout,
		Stderr: m.Stderr,
	}
}

// Run executes the command.
func (cmd *InfoCommand) Run(args ...string) error {
	// Parse flags.
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	help := fs.Bool("h", false, "")
	if err := fs.Parse(args); err != nil {
		return err
	} else if *help {
		fmt.Fprintln(cmd.Stderr, cmd.Usage())
		return ErrUsage
	}

	// Require database path.
	//path := fs.Arg(0)
	//if path == "" {
	//	return ErrPathRequired
	//} else if _, err := os.Stat(path); os.IsNotExist(err) {
	//	return ErrFileNotFound
	//}

	// Open the database.
	//db, err := bolt.Open(path, 0666, nil)
	//if err != nil {
	//	return err
	//}
	//defer db.Close()
	//
	//// Print basic database info.
	//info := db.Info()
	//fmt.Fprintf(cmd.Stdout, "Page Size: %d\n", info.PageSize)

	return nil
}

// Usage returns the help message.
func (cmd *InfoCommand) Usage() string {
	return strings.TrimLeft(`
usage: bolt info PATH
Info prints basic information about the Bolt database at PATH.
`, "\n")
}

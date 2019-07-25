package main

import (
	"os"
	userconsole "playgo/awesomelib/clis/kingpin/modules/user/console"
	storeconsole "playgo/awesomelib/clis/kingpin/modules/store/console"
)
import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	name    = kingpin.Arg("name", "Name of user.").Required().String()
)

func configureLsCommand(app *kingpin.Application) {
	c := &userconsole.LsCommand{}
	ls := app.Command("ls", "List files.").Action(c.Run)
	ls.Flag("all", "List all files.").Short('a').BoolVar(&c.All)

	c2 := &storeconsole.LsCommand{}
	ls2 := app.Command("ls-store", "List files.").Action(c2.Run)
	ls2.Flag("all", "List all stores.").Short('a').BoolVar(&c2.All)

}

func main() {
	//kingpin.Parse()
	//fmt.Printf("%v, %s\n", *verbose, *name)

	app := kingpin.New("modular", "My modular application.")
	configureLsCommand(app)

	userconsole.ConfigureCommands(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
package main

import (
	"github.com/lalamove/konfig"
	_ "github.com/lalamove/konfig"
	"github.com/lalamove/konfig/loader/klfile"
	"github.com/lalamove/konfig/parser/kpjson"
	"github.com/micro/go-micro/util/log"
)
var configFiles = []klfile.File{
	{
		Path:   "./config.json",
		Parser: kpjson.Parser,
	},
}

func init() {
	konfig.Init(konfig.DefaultConfig())
}

func main() {
	// load from json file
	konfig.RegisterLoaderWatcher(
		klfile.New(&klfile.Config{
			Files: configFiles,
			Watch: true,
		}),
		// optionally you can pass config hooks to run when a file is changed
		func(c konfig.Store) error {
			return nil
		},
	)

	if err := konfig.LoadWatch(); err != nil {
		log.Fatal(err)
	}

	// retrieve value from config file
	debug := konfig.Bool("debug")
	print(debug)

	// select{}
}
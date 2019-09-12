package main

import (
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	printer.Fprint(os.Stdout, fs, f)
}

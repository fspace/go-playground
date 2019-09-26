package main

import (
	"github.com/nathanfaucett/inflect"
)

func main() {
	// also PluralizeLocale(string, locale)
	s := inflect.Pluralize("string") // return "strings"
	println(s)
}

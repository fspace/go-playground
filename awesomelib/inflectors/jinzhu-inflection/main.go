package main

import "github.com/jinzhu/inflection"

func main() {
	println(inflection.Plural("FancyPerson")) // => "FancyPeople"
}

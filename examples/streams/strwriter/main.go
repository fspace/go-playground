package main

import "strings"

func main() {
	b := strings.Builder{}
	b.WriteString("One")
	c := b
	c.WriteString("Hey!") // panic: strings: illegal use of non-zero Builder copied by value
}

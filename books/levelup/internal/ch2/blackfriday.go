package ch2

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func Blackfriday() {
	markdown := []byte(`
# This is a header
* and
* this
* is
* a
* list
`)
	html := blackfriday.MarkdownBasic(markdown)
	fmt.Println(string(html))
}

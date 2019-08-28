package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(nil)
	b.WriteString("One")
	s1 := b.String()
	b.WriteString("Two")
	s2 := b.String()
	b.Reset()
	b.WriteString("Hey!") // does not change s1 and s2
	s3 := b.String()
	fmt.Println(s1, s2, s3) //
}

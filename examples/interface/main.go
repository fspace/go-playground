package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

type Stringer interface {
String() string
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}

package lib

import "fmt"

type A struct{}

func (A) Method() { fmt.Println("Hello, playground from method of struct") }

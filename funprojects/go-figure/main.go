package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}

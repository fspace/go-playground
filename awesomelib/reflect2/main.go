package main

import (
	"fmt"
	"github.com/modern-go/reflect2"
)

func demo1() {
	type widget struct {
	}
	w := widget{}
	tpy := reflect2.TypeOf(w)

	fmt.Println("...")
	fmt.Printf("%#v", tpy)
}
func main() {

	i := 1
	tpy := reflect2.TypeOf(i)

	fmt.Println("type of i is :", tpy)

	//
	demo1()

}

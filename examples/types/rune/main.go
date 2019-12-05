package main

import "fmt"

func main() {
	demo1()
	zeroVal()
}

func demo1() {
	r := 'a'
	fmt.Printf("%v , %T \n", r, r)

	var r2 rune = 'a'
	fmt.Printf("%v , %T \n", r2, r2)
}

func zeroVal() {
	var r rune
	fmt.Printf("%v , %T \n", r, r)

}

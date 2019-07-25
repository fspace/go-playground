package main

import "fmt"

func testPtr(num *int){
	*num = 8
}

func main() {
	num := 1
	fmt.Println("current num is :", num)
	testPtr(&num)
	fmt.Println("now num is :",num)
}

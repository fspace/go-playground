package main

import "fmt"

func getSumAndSub(a, b int)(sum, sub int)  {
	sum = a+b
	sub = a - b
	return
}

func main() {
	sum , sub := getSumAndSub(3 , 2)
	fmt.Println("sum is ", sum)
	fmt.Println("sub is ", sub)
}

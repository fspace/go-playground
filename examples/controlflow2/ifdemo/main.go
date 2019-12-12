package main

import (
	"fmt"
	"math"
)

func main() {
	simple()
}

func simple() {
	if true {
		fmt.Println("This test is a true")
	}
}

func basic() {
	number := 50
	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must between 1 and 100!")
	}
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("too high")
	}
	if guess == number {
		fmt.Println("You got it !")
	}
	if true && boolTest() {
		fmt.Println("Short circuit expr")
	}
	if true || boolTest() {
		fmt.Println("this is never be executed!")
	}
}

func boolTest() bool {
	fmt.Println("return true")
	return true
}

func usecaseNumberCompare() {
	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

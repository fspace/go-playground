package main

import "fmt"

func ifStmt() {
	a := 8
	if r := a % 10; r != 0 { // 伴有短声明
		if r > 5 {
			a -= r
		} else if r < 5 {
			a += 10 - r
		}
	} else {
		a /= 10
	}
}

func switchStmt() {
	tier := 3
	age := 20
	switch tier { // switch statement
	case 1: // case statement
		fmt.Println("T-shirt")
		if age < 18 {
			break // exits the switch block
		}
		fallthrough // executes the next case
	case 2:
		fmt.Println("Mug")
		fallthrough // executes the next case
	case 3:
		fmt.Println("Sticker pack")
	default: // executed if no case is satisfied
		fmt.Println("no reward")
	}
}

// The for statement is the only looping statement in Go.
func forStmt() {
	condition := false
	for { // infinite loop
		if condition {
			break // exit the loop
		}
	}
	i := 10
	for i < 0 { // loop with condition
		if condition {
			continue // skip current iteration and execute next
		}
	}

	for i := 0; i < 10; i++ { // loop with declaration, condition and operation
	}
}

func forStmt2() {
	a := 5
label:
	for i := a; i < a+2; i++ {
		switch i % 3 {
		case 0:
			fmt.Println("divisible by 3")
			break label // this break the outer for loop
		default:
			fmt.Println("not divisible by 3")
		}
	}
}

func main() {
	ifStmt()
	switchStmt()
}

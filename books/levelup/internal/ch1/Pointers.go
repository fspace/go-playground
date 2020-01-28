package ch1

import "fmt"

func PassingByReference() {
	fruit := "banana"
	giveMePear(&fruit)
	fmt.Println(fruit) // Outputs: pear
}
func giveMePear(fruit *string) {
	*fruit = "pear"
}

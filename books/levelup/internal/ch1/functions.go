package ch1

import "fmt"

func DefineFunc() {
	// While you can return more than two values, it’s not common to see this

	noParamsNoReturn()
	_ = twoParamsOneReturn(1, "yes")
	twoSameTypedParams("hello", "go")

	a, b := oneParamTwoReturns(3)
	fmt.Println(a) // Outputs "Int: 3"
	fmt.Println(b) // Outputs "4"
}

func noParamsNoReturn() {
	fmt.Println("I’m not really doing much!")
}
func twoParamsOneReturn(myInt int, myString string) string {
	return fmt.Sprintf("myInt: %d, myString: %s", myInt, myString)
}
func oneParamTwoReturns(myInt int) (string, int) {
	/*
		A lot of functions will return two values: one for the actual response from the
	function and a second to indicate whether or not an error has occurred. This is a
	very common practice in Go and forms the basis for handling errors,

	*/
	return fmt.Sprintf("Int: %d", myInt), myInt + 1
}
func twoSameTypedParams(myString1, myString2 string) {
	fmt.Println("String 1", myString1)
	fmt.Println("String 2", myString2)
}

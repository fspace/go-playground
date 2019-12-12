package main

import "fmt"

func main() {
	basic()
	initializer()
	demo3()
	demo4()

	typeSwitch()
	typeSwitch2()
}
func initializer() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one , five or ten")
	default:
		fmt.Println("another number")
	}
}
func basic() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
func demo2() {
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one , five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("not ...")

	}
}
func demo3() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		// break // 没必要哦
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater then twenty")
	}
}
func demo4() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // 下穿
	case i > 20: // 因为fallthrough的缘故 假的也执行！
		fmt.Println("greater than   to twenty")
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater then twenty")
	}
}

func typeSwitch() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")

	}
}
func typeSwitch2() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		fmt.Println("this will be printed too!")
		break // 下面的就可以跳出执行了！
		fmt.Println("this will never be  printed !")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")

	}
}

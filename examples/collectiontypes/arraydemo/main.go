package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("current wd is : ", wd)

	problem()

	creationDemo()
	workingWithArrays()
	buildinOps()

	creationDemo2()
	usecase1()
	usecase2()
}

func problem() {
	grade1 := 97
	grade2 := 93
	grade3 := 85
	fmt.Printf("Grades: %v , %v, %v \n", grade1, grade2, grade3)
}

func creationDemo() {
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v \n", grades)
	// 数量推导
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades2: %v \n", grades2)

	// 零值
	var students [3]string
	fmt.Printf("Students: %v \n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v \n", students)
}
func creationDemo2() {
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println(identityMatrix)
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)
}
func buildinOps() {
	var students [3]string
	fmt.Printf("Students: %v \n", students)
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "Arnold"
	fmt.Printf("Number of Students : %v \n", len(students))
}

func workingWithArrays() {
	// 零值
	var students [3]string
	fmt.Printf("Students: %v \n", students)
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "Arnold"
	fmt.Printf("Students #1: %v \n", students[1])
}
func usecase1() {
	fmt.Println(">--- copy ----------------")

	a := [...]int{1, 2, 3}
	b := a // 值拷贝
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
func usecase2() {
	fmt.Println(">--- copy with pointer ----------------")

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}

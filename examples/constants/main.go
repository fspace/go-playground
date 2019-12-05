package main

import "fmt"

const MAX_Xx = 100                    // integer
const MAX_Xx2 = 20.99                 // float
const EVENT_USER_ADDED = "user-added" // string

const STATUS_ON = true // boolean
const STATUS_OFF = false

// const RGBS  = [...]int{255} // wrong！

const PiApprox = 3.14

//var PiInt int = PiApprox  // 3, converted to integer
//var Pi float64 = PiApprox // is a float
//
//type MyString string
//
//const Greeting = "Hello!"
//
//var s1 string = Greeting   // is a string
//var s2 MyString = Greeting // string is converted to MyString

const a int16 = 27

func main() {

	constDemo()
	shadow()

	typeInfer() // 编译器类型推断

	iotaDemo()
	iotaDemo2()
	iotaDemo3()
	iotaDemo4()
	iotaDemo5()
}

func constDemo() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
}

// 外部同名常量被遮盖
func shadow() {
	const a int16 = 3
	fmt.Printf("%v, %T\n", a, a)
}

func typeInfer() {
	const a = 42 // a 类型受后面运算类型影响！
	var b int16 = 27
	fmt.Printf("%v, %T\n", a+b, a+b)
}

const n = iota // 常量 数字型计数器

func iotaDemo() {
	fmt.Printf("%v, %T\n", n, n)
}

const (
	c1 = iota
	c2 = iota
	c3 // 编译器推导
	c4 // 编译器推导
)

const (
	STATUS_A = iota
	STATUS_B = iota
)

func iotaDemo2() {
	fmt.Println("iota-demo2 ===============================")

	fmt.Printf("%v, %T\n ", c1, c1)
	fmt.Printf("%v, %T\n ", c2, c2)
	fmt.Printf("%v, %T\n ", c3, c3)
	fmt.Printf("%v, %T\n ", c4, c4)

	// 另一个常量块
	fmt.Printf("%v, %T\n ", STATUS_A, STATUS_A)
	fmt.Printf("%v, %T\n ", STATUS_B, STATUS_B)
}

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func iotaDemo3() {
	fmt.Printf("%v\n", catSpecialist)
}

const (
	_  = iota // 忽略第一个0值 通过赋值给一个下划线
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func iotaDemo4() {
	fileSize := 400000000.
	fmt.Printf("%.2f GB", fileSize/GB)
}

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func iotaDemo5() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("\n %b\n", roles)

	fmt.Printf("IS Admin? %v \n", isAdmin&roles == isAdmin)
	fmt.Printf("IS HQ? %v \n", isHeadquarters&roles == isHeadquarters)
}

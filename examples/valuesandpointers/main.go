package main

import "fmt"

// 基于值 拷贝！

type A int

func (a A) Foo() {
	a++
	fmt.Println("foo", a)
}

func main() {
	var a A
	fmt.Println("before", a)
	a.Foo()
	a.Foo()
	a.Foo()
	fmt.Println("after", a)
	// ==============================================
	u := User{Name: "Yiqing", Age: 18}
	fmt.Println(u.Name, "is now ", u.Age)
	u.Birthday()
	fmt.Println(u.Name, "is now ", u.Age)
	// ------------------
	u.Birthday2()
	fmt.Println(u.Name, "is now ", u.Age)
	// .................

	u.Birthday3()
	fmt.Println(u.Name, "is now ", u.Age)

	// ==============================
	// 天生传引用的类型 map slice channel 这些类型内部设计就是一个含有指针的结构
}

type User struct {
	Name string
	Age  int
}

func (u User) Birthday() {
	u.Age++
	fmt.Println(u.Name, "turns ", u.Age)
}
func (u *User) Birthday2() {
	u.Age++
	fmt.Println(u.Name, "turns ", u.Age)
}
func (u *User) Birthday3() {
	*u = User{Name: u.Name, Age: u.Age + 1} // 整个替换掉了
	fmt.Println(u.Name, "turns ", u.Age)
}

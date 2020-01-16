package models

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) DoSomething() {
	fmt.Println(p.Name, "is doing something! ")
}

func (p *Person) DoSomething2() {
	fmt.Println(p.Name, "is doing something! ")
}

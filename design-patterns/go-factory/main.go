package main
// @see https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/
type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson0(name string, age int) Person {
	return Person{
		Name: name,
		Age: age
	}
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age
	}
}

func main() {
	
}

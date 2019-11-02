-  Use factory functions to ensure that new instances structs are constructed with the required arguments

-  Use interface factories when you want to use multiple implementations interchangeably

-  Use factory methods to make factories that return inter-related instances of different structs or interfaces

-  Factory functions are useful when you want a factory to make specialized instances of a single struct.
Factory functions
We can use factory functions to make factories with some defaults built in. For example, if we want to make a 
“baby factory” and a “teenager factory” instead of a generic “person factory”, we use a generator function:
~~~go
type Person struct {
	name string
	age int
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			name: name,
			age: age,
		}
	}
}

// ---------------------------------
newBaby := NewPersonFactory(1)
baby := newBaby("john")

newTeenager := NewPersonFactory(16)
teen := newTeenager("jill")
~~~

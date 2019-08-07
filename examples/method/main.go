package main

import (
	"fmt"
	"playgo/examples/method/lib"
)

type A int

func (a A) Foo() { fmt.Println("Foo method of type A , and a is :", int(a)) }

// ========================================================================================= |
type ErrKey string

func (e ErrKey) Error() string {
	//return fmt.Errorf("key %q not found", e)
	return fmt.Sprintf("key %q not found", e)
}

var _ error = ErrKey("") // 确保类型实现了接口  可以在编译期发现问题
// ========================================================================================= |

func main() {
	a := A(1)
	a.Foo()  // Call the method on an instance of the type
	A.Foo(a) // Call the method on the type and passing an instance as argument

	ms := lib.MyStruct{}
	ms.Foo()

}

package main

import (
	"fmt"
)

// ## 类型别名和原类型完全一样，只不过是另一种叫法而已

type AddSum = int //给 int 取一个别名叫 AddSum

//func (this AddSum)printMyType() {
//	fmt.Printf("AddSum underline type is %T \n", this)
//}

type NewInt int //定义 NewInt 的类型为 Int
func (this NewInt) printMyType() {
	fmt.Printf("AddSum underline type is %T \n", this)
}

func main() {
	var a AddSum
	fmt.Printf("%T\n", a)
	var a2 NewInt
	fmt.Printf("%T\n", a2)

	a2.printMyType()
	// 结构体别名测试

	obj := SomeStruct{}
	var objWrapper SomeStructAlias = obj
	objWrapper.printMyType()
	obj.printMyType()
}

type SomeStruct struct {
}

// 类型别名
type SomeStructAlias = SomeStruct

// printMyType
// 定义在类型别名 SomeStructAlias 上 但原始类型SomeStruct 也具有了方法printMyType!
func (this SomeStructAlias) printMyType() {
	fmt.Printf("SomeStructAlias underline type is %T \n", this)
}

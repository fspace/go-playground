package main

import (
	"fmt"
)

// ## 类型别名和原类型完全一样，只不过是另一种叫法而已
// 完全一样(identical types)意味着这两种类型的数据可以互相赋值，而类型定义要和原始类型赋值的时候需要类型转换(ConversionT(x))。

// 进一步想，这样是不是我们可以为其它库中的类型增加新的方法了， 比如为标准库的time.Time增加一个滴答方法
// 答案是:NO , 编译的时候会报错:cannot define new methods on non-local type time.Time。

type AddSum = int //给 int 取一个别名叫 AddSum

// NOTE  这里就是 不能为其他包中的类型定义方法！
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

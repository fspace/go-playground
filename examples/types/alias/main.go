package main

import (
	"fmt"
	"strings"
)

// @see https://colobu.com/2017/06/26/learn-go-type-aliases/  rpcx作者：https://github.com/smallnest/rpcx
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

	println(strings.Repeat("=", 80))
	role1()
	role2()
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

// ===============================================================================  |
//					## DCI 架构
// - yii中 组件类 可以动态绑定行为|或者设计期声明式的绑定行为 这样对象就有了某种角色
// dci中 把一次交互 认为是舞台上的一个场景片段   先挑选不同的对象  然后让他们承担不同的角色 然后按照角色剧本表演就好了
// 考量后  还是内嵌类型可能最贴近了  因为内嵌可以添加其他属性  这个在yii中Behavior也是这样 可以附加其他属性和方法
// 如果是新类型声明 可能不具备扩展属性的能力         比如我要扮演包公 你得给我头上弄个月亮吧！

// TODO  要实现DCI架构  将类型跟角色绑定 考虑两种可能 1. 内嵌  2. 类型
type User struct {
	Name string
	Age  int
}

type Eater struct {
	User
}

func (this Eater) Eat() {
	fmt.Printf("I am %s and my age is %d \n ", this.Name, this.Age)
	fmt.Println("I'm Eating something !")
}

func role1() {
	u := User{
		Name: "yiqing",
		Age:  18,
	}
	etr := Eater{User: u} // 相当于角色绑定
	etr.Eat()
}

type Eater2 User

func (this Eater2) Eat() {
	fmt.Printf("I am %s and my age is %d \n ", this.Name, this.Age)
	fmt.Println("I'm Eating something !")
}

func role2() {
	u := User{
		Name: "yiqing",
		Age:  18,
	}
	etr := Eater2(u)
	etr.Eat()
}

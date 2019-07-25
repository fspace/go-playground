package lib

import (
	"fmt"
	"strconv"
)

/**
每种类型对应的方法必须和类型的定义在同一个包
中，因此是无法给 int 这类内置类型添加方法的（因为方法的定义和类型的定义
不在一个包中）
 */

type MyStruct struct {

}

func (ms MyStruct)Foo()  {
	fmt.Printf("this is a method of struct : %T \n", ms)
	someInt := 10
	mi := myInt(someInt)
	fmt.Println(mi.ToString())
	fmt.Println(myInt2string(&mi))
	fmt.Printf( "func type is %T", myInt2string)

	mi.Incr(10)
	fmt.Println(mi.ToString())
	// 另一种递增调用方法
	myIntIncr(&mi,10)
	fmt.Println(mi)
	fmt.Printf( "func type of myIntIncr is %T", myIntIncr)
}

type myInt int // 类型声明

func (mi myInt)ToString() string  {
	fmt.Println("call toString of MyInt")
	return strconv.FormatInt(int64(mi),10)
}
func (mi *myInt)Incr(delta int)  {
	*mi = *mi + myInt(delta)
}

type myInt2 = int  // 类型别名
/*
// 不能为内置类型声明方法哦！
func (mi myInt2)ToString() string  {
	fmt.Println("call toString of MyInt")
	return strconv.FormatInt(int64(mi),10)
}
*/
// 方法 函数  转换！
// 方法表达式
var myInt2string = (*myInt).ToString
var myIntIncr = (*myInt).Incr
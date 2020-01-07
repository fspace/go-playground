package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	//basic2()
	//ptrArray()
	realmain()
}
func realmain() {
	app := cli.App("ptr-demo", "pointer demo")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("struct", "ptr about struct", cli.ActionCommand(ptrStruct))

	app.Command("arr-slice", "数组 切片 赋值后 修改的影响", cli.ActionCommand(arrayAndSlice))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)
}

func basic() {
	a := 42
	fmt.Println(a)
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
}
func basic2() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)

	fmt.Println(&a, b)
	// dereference
	fmt.Println(a, *b)

	// 修改
	a = 27
	fmt.Println(a, *b)
	fmt.Printf("addr: %p", b)
}

func ptrArray() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1] // 不允许加减哦  真想这么搞 可以看看unsafe包的东东！
	// 数组元素指针位置临近 4|8 字节？
	fmt.Printf("%v %p %p\n", a, b, c)
}

type myStruct struct {
	foo int
}

func ptrStruct() {
	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{42}
	fmt.Println(ms)

	ms2 := new(myStruct) // 此种形式不能指定初始值！
	fmt.Println(ms2)
	// 解引用 dereference
	(*ms2).foo = 42
	fmt.Println((*ms).foo)
	// 多余的
	ms2.foo = 55
	fmt.Println(ms.foo)
}

func arrayAndSlice() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)

	a2 := []int{1, 2, 3} // 换为了切片
	b2 := a2
	fmt.Println(a2, b2)
	a2[1] = 42
	fmt.Println(a2, b2)
}

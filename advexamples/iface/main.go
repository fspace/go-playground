package main

import (
	"playgo/advexamples/iface/internal"
	"playgo/advexamples/iface/internal/impl1"
	"playgo/advexamples/iface/internal/impl2"
	nil2 "playgo/advexamples/iface/internal/nil"
)

func main() {
	demo1()

	demo2()

	nilDemo()
}

func demo1() {
	type Test struct {
	}

	v := Test{}

	Print(v)
}
func Print(v interface{}) {
	println(v)
}

func demo2() {
	var d internal.Duck = impl1.Cat{}
	var d2 internal.Duck = &impl2.Cat{}

	var d3 internal.Duck = &impl1.Cat{} // 也能通过编译！

	//d.Walk()
	//d.Quack()
	//d2.Walk()
	//d2.Quack()
	//d3.Walk()
	//d3.Quack()
	internal.DuckQuack(d, d2, d3)
	internal.DuckWalk(d, d2, d3)
}

func nilDemo() {
	var s *nil2.TestStruct
	nil2.NilOrNot(s)
}

func demo3() {
	// 用 结构体实现接口方法 impl1 也可以的！
	var c internal.Duck = &impl2.Cat{}
	switch c.(type) {
	case *impl2.Cat:
		cat := c.(*impl2.Cat)
		cat.Quack()
	}
}

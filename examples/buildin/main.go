package main

import (
	"fmt"
	"os"
)

func deferExample() error {
	f, err := os.Open("config.txt")
	if err != nil {
		return err
	}
	defer f.Close() // 无论如何总是会关闭的   defer的列表会以栈的形式执行--FILO 先进后出
	// 用f做一些操作
	return nil
}

// 主要是给值类型分配空间
func newDemo1() {
	n := 1
	fmt.Printf("n 类型 %T ; n 的值：%v ; n的地址是：%v  \n", n, n, &n)

	n2 := new(int)
	fmt.Printf("n2 类型 %T ; n2 的值：%v ; n2 的地址是：%v \n", n2, n2, &n2)
	fmt.Printf("n2 所指指针所指的值是：%v \n", *n2)
}

// 给引用类型分配空间 主要是 channel map slice
func makeDemo1() {

}

func main() {
	newDemo1()
}

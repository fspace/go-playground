package main

import "fmt"

func main() {
	creationDemo()
	demo2()

	buildin1()
	buildin2()
	usecasePop()
}
func creationDemo() {
	a := []int{1, 2, 3}
	fmt.Println(a)

	// builtin funcs
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// assignment
	b := a // 共享同一底层结构
	b[1] = 5
	fmt.Println(b)
	fmt.Println(a)

}

func demo2() {
	// 底层同一数组
	// a := []int{1,2,3,4,5,6,7,8,9,10}
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 数组 | 切片 都可以哦 换成上面的也行 下面的操作都是一样结果
	b := a[:]                                    // 全元素切片
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	a[5] = 42 // 写元素  一改全改
	fmt.Println(a, b, c, d, e)
}

func buildin1() {
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Len: %v\n", len(a))
	fmt.Printf("Cap: %v\n", cap(a))

	a2 := make([]int, 3, 100)
	arrayInfo(a2)
}
func buildin2() {
	a := []int{}
	arrayInfo(a)
	a = append(a, 1)
	arrayInfo(a)

	a = append(a, 2, 3, 4, 5) // 可变参数
	arrayInfo(a)

	// 多数组合并
	a = append(a, []int{6, 7, 8}...) // 元素展开...
	arrayInfo(a)
}
func arrayInfo(s []int) {
	fmt.Println("array|slice : ", s)
	fmt.Printf("length: %v \n", len(s))
	fmt.Printf("Capacity: %v \n", cap(s))
}

func usecasePop() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:]
	arrayInfo(b)
	c := a[:len(a)-1]
	arrayInfo(c)
	// 中间移除
	d := append(a[:2], a[3:]...)
	arrayInfo(d)
	arrayInfo(a) // a 变得很奇怪了
}

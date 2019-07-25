package main

import "fmt"

func defSlice(){
	var(
		a  []int  // nil 切片 和nil相等， 一般用来表示一个不存在的切片
		b = []int{} // 空切片， 和nil不相等 一般用来表示一个空集合
		c = []int{1, 2, 3} // 有3个元素的切片, len和cap都为3
		d = c[:2] // 有2个元素的切片, len为2, cap为3
		e = c[0:2:cap(c)] // 有2个元素的切片, len为2, cap为3
		f = c[:0] // 有0个元素的切片, len为0, cap为3
		g = make([]int, 3) // 有3个元素的切片, len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
		i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	)
	fmt.Println(
		a,
		b,
		c,
		d,
		e,
		f,
		g,
		h,
		i,		)

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
}

func opAppend()  {
	var a []int
	a = append(a, 1) // 追加1个元素
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包

	fmt.Printf("a is %#v", a)
}

func  prepend()  {
	var a = []int{1,2,3}
	a = append([]int{0}, a ...)
	a = append([]int{-3, -2, -1}, a...)
	fmt.Printf("%v", a)
}
/**
Inpend 从中间插入元素

Copy 的另一个语义 就是“移动"
 */
func inpend()  {
	var a []int
	i := 0
	a = append(a[:i], append([]int{1},a[:i]...)...) // 第i个位置插入一个元素

	a = append(a[:i], append([]int{1,2,3}, a[:i]...)...)
	// 每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内
	// 容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 。
	// 可以用 copy 和 append 组合可以避免创建中间的临时切片，同样是完成添加元
	// 素的操作：
	a = append(a, 0) // 切片扩展一个空间
	copy(a[i+1:] , a[i:])
	a[i] = 2

	fmt.Printf("\n %#v",a)
	// 中间插入切片
	x := []int{10,11}
	a = append(a, x...) // 为x切片扩展足够的空间
	copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
	copy(a[i:], x) // 复制新添加的切片
	fmt.Println(a)
}
// del
func delElement()  {
	var N  = 1
	// 头 中 尾部 删除   三种情况 末尾删除最快
	a := []int{1,2,3,4,5,6,7,8}
	a = a[:len(a)-1] // 删除尾部一个元素
	a = a[:len(a)-N] // 删除尾部N个元素
	fmt.Println(a)

	// 删除开头元素  直接移动数据指针：
	a = a[1:] // 删除开头一个元素
	fmt.Println(a)
	a = a[N:] // 删除开头 N 个元素
	fmt.Println(a)

	a = append(a[:0],a[1:]...) // 删除开头的一个元素
	fmt.Println(a)
	a = append(a[:0],a[N:]...) // 删除开头N个元素
	fmt.Println(a)

	// COPY 完成删除元素：
	a = []int{1,2,3}
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	a = a[:copy(a, a[N:])] // 删除开头的N个元素

	a = []int{1,2,3,4,5,6,7,8,9}
	a = append(a[:i], a[i+1:]...)

	a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	a = a[:i + copy(a[i:],a[i+N:])] // 删除中间N个元素

}

func Filter(s []byte, fn func(x byte) bool) []byte{
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}
func main() {
	opAppend()
	prepend()
	inpend() // 中间插入
     // 删除
     delElement()
}

package main

import "fmt"

func ops() {
	var a int
	a = 10
	fmt.Println("a = ", a)
	a2 := 20
	fmt.Println("a2 = ", a2)
	fmt.Println("a == a2 :", a == a2)
	if a != a2 {
		fmt.Println("a != a2 is true")
	}

	fmt.Println("a1 sum a is :", a+a2)
	fmt.Println("a minus a2 is :", a-a2)
	fmt.Println("a times a2 is :", a*a2)
	fmt.Println("a divide a2 is :", a/a2)

	fmt.Println("3 modulo 2 is :", 3%2)

	fmt.Println("  1 & 0 is:", 1&0)
	fmt.Println("  1 &^ 0 is:", 1&^0)
	fmt.Println("  1 << 1 is:", 1<<1)
	fmt.Println("  1 << 2 is:", 1<<2)
	fmt.Println("  1 << 10 is:", 1<<10)
	fmt.Println("  1 << 20 is:", 1<<20)

	fmt.Println("  4 >> 1 is:", 4>>1) // 4 / 2
	fmt.Println("  4 >> 2 is:", 4>>2) // 4 / 4

	fmt.Println("  false && true is:", false && true)
	fmt.Println("  true && true is:", true && true)
	fmt.Println("  false && false is:", !true && !true)

	fmt.Println("  true || false is:", true || false)

	fmt.Println("  not true is:", !true)

	var i = 2
	ri := &i // 引用 Returns the pointer to a variable
	*ri = 3  // 解引用 Dereference  Returns the content of a pointer
	fmt.Println("now i is :", i)
}

func ops2() {
	//var c1 chan int // 这样搞就是死锁！  未分配空间？
	c1 := make(chan int)
	go func() {
		fmt.Println(<-c1)

	}()
	c1 <- 1
}

func main() {
	ops()
	// ops2() // 死锁！
	ops2()
}

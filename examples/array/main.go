package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

var a [3]int
var b = [...]int{1, 2, 3}
var c = [...]int{2: 3, 1: 2}
var d = [...]int{1, 2, 4: 5, 6}

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	bigArr()

	typeArray()

	usageOfEmptyArray()
}

func bigArr() {
	var a = [...]int{1, 2, 3, 40: 40, 41}
	var b = &a // b 是指向数组的指针
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1]) // // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b {
		fmt.Println(i, v)
	}
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
	// 我们可以用 fmt.Printf 函数提供的 %T 或 %#v 谓词语法来打印数组的类型和详
	// 细信息：
	fmt.Printf("b: %T\n", b)  // b: [3]int
	fmt.Printf("b: %#v\n", b) // b: [3]int{1, 2, 3
}

func typeArray() {
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}
	fmt.Printf("s-array1: %v \n", s1)
	fmt.Printf("s-array1: %v \n", s2)
	fmt.Printf("s-array1: %v \n", s3)

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line1, line2, line3)

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	fmt.Println(decoder1, decoder2)

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}
	fmt.Println(unknown1, unknown2)
	// 管道数组
	var chanList = [2]chan int{}
	fmt.Println(chanList)

	var d [0]int       // 定义一个长度为0的数组
	var e = [0]int{}   // 定义一个长度为0的数组
	var f = [...]int{} // 定义一个长度为0的数组
	fmt.Println(d, e, f)

}

/**
长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于
强调某种特有类型的操作时避免分配额外的内存空间，比如用于管道的同步操作：
*/
func usageOfEmptyArray() {
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1
	/**
		对于这种场景，我们用空数组来作为管道类型可以减少管道元素
	赋值时的开销。当然一般更倾向于用无类型的匿名结构体代替：
	*/
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2
}

// ==============================================================
// https://tutorialedge.net/golang/go-complex-types-tutorial/
func basic() {
	// declaring an empty array of strings
	// var days []string

	// declaring an array with elements
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println(days[0]) // prints 'monday'
	fmt.Println(days[5]) // prints 'saturday'

	// slice
	weekdays := days[0:5]
	fmt.Println(weekdays)

}

// ==============================================================

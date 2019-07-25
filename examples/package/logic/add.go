package logic

import "fmt"

const PI  = 3.141592

var  x = 0

func init() {
	fmt.Println("logic:init is called! ")
	fmt.Printf("PI is %v \n",PI)
	fmt.Printf("var x is %v \n ",x)
	fmt.Println()

	// init中的 携程运行在main之后
	go func() {
		fmt.Println("hi  am a goroutine from package:logic")
	}()
}



func Add(a , b int) int {
	return a + b
}

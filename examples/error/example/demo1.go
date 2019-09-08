package example

import "fmt"

func Foo1() {
	n1 := 10
	n2 := 0
	rslt := n1 / n2
	fmt.Println("hi you can't go here! ", rslt)
}

// defer recover panic  这种机制可以捕获异常
// 本例是内部自己处理了panic 而不是抛给调用者
func Foo2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err occurred in foo2: ", err)
		}
	}()
	n1 := 10
	n2 := 0
	rslt := n1 / n2
	fmt.Println("hi you can't go here! ", rslt)
}

func Foo3() (err error) {
	defer func() {
		er := recover()
		if er != nil {
			fmt.Println("err occurred in foo3: ", er)
			// 可以在defer中做返回值串改
			// 转型赋值给最终的返回值
			err = er.(error)
			fmt.Println("发个消息给管理员！")
		}
	}()
	n1 := 10
	n2 := 0
	rslt := n1 / n2
	fmt.Println("hi you can't go here! ", rslt)
	return nil
}

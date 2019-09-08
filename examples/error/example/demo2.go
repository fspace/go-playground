package example

import (
	"fmt"
	"github.com/pkg/errors"
)
import _ "github.com/pkg/errors"

func Foo4() (err error) {
	defer func() {
		er := recover()
		if er != nil {
			// fmt.Println("err occurred in foo4: ", er)
			// 可以在defer中做返回值串改
			// 转型赋值给最终的返回值
			err = er.(error)
			fmt.Println("发个消息给管理员！")
			err = errors.Wrap(err, "Foo4 error happened！")
		}
	}()
	n1 := 10
	n2 := 0
	rslt := n1 / n2
	fmt.Println("hi you can't go here! ", rslt)
	return nil
}

func CreateError() error {
	return errors.New("this is a customm-error")
}

// CreateError2 测试触发异常 可传递任意类型  recover捕获的不一定是error类型 就是panic抛出的东西
func CreateError2() error {
	defer func() {
		if err := recover(); err != nil {
			// 看下err 类型
			fmt.Printf("panic target is %v and the type is %T : \n", err, err)
		}
	}()
	panic(2)
	return nil
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取...
		return nil
	} else {
		return errors.New(fmt.Sprintf("文件名称[%s]不正确...", name))
	}
}
func Exec1() {
	err := readConf("config.ini2")
	if err != nil {
		panic(err)
	}
	fmt.Println("continue ...")
}

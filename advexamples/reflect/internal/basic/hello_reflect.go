package basic

import (
	"fmt"
	"reflect"
)

// https://draveness.me/golang/basic/golang-reflect.html
func HelloReflect() {
	author := "yiqing"
	fmt.Println("TypeOF author: ", reflect.TypeOf(author))
	fmt.Println("ValueOf author: ", reflect.ValueOf(author))
}

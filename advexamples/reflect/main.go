package main

// https://gfw.go101.org/article/reflection.html
import "fmt"
import "reflect"

type F func(string, int) bool

func (f F) Validate(s string) bool {
	return f(s, 32)
}

func main() {
	var x struct {
		n int
		f F
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())     // struct
	fmt.Println(tx.NumField()) // 2
	tf := tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1
	fmt.Println(tf.NumMethod())          // 1
	ts, ti, tb := tf.In(0), tf.In(1), tf.Out(0)
	fmt.Println(ts.Kind(), ti.Kind(), tb.Kind()) // string int bool
}

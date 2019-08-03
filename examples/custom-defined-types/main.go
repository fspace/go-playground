package main

// https://blog.altoros.com/golang-internals-part-2-diving-into-the-go-compiler.html
type I interface {
	DoSomeWork()
}

type T struct {
}

func (t *T) DoSomeWork() {

}

func main() {
	t := &T{}
	i := I(t)
	println(i)
}

type Message string
type Counter int
type Number float32
type Success bool

// ============================================
// 组合自定义类型
type StringDuo [2]string
type News chan string
type Score map[string]int
type IntPtr *int
type Transform func(string) string
type Result struct {
	A, B int
}

// 合起来用
type Broadcast Message
type Timer Counter
type News chan Message

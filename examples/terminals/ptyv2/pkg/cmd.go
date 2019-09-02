package pkg

import "io"

// All of the information about each command can be self-contained in a structure.
type Cmd struct {
	Name string // the cmd name
	Help string // a description string
	// 结构体 可以看做成员变量集合 + 方法集合  方法如果想再灵活一点 可以变为函数类型的变量
	// 由外部设置 这样灵活度就提升了 这是一个技巧
	// 把结构体方法 提升为成员变量 此后就可以变为依赖注入啦 然后可以外部注入模拟的实现 测试时会很方便
	Action func(w io.Writer, args ...string) bool
}

func (c Cmd) Match(s string) bool {
	return c.Name == s
}

func (c Cmd) Run(w io.Writer, args ...string) bool {
	return c.Action(w, args...)
}

package main

import (
	"github.com/grsmv/inflect"
)

// 可以看下测试比较有意思 https://github.com/grsmv/inflect/blob/master/inflect_test.go#L294

func main() {
	println(inflect.Parameterize("Trailing bad characters!@#")) //:          "trailing-bad-characters",
}

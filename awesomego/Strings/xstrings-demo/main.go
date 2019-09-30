package main

import (
	"fmt"
	"github.com/huandu/xstrings"
)

// 该库还是有很多类似inflection 的功能
func main() {
	fmt.Println(xstrings.ToCamelCase("yes_you_see"))
}

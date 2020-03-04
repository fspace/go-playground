package main

import "fmt"

// source: https://stackoverflow.com/questions/24790175/when-is-the-init-function-run
var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func init() {
	fmt.Println("Called second in order of declaration")
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}

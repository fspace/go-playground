package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://www.baidu.com/").End()

	fmt.Println("errs:", errs)
	fmt.Println("resp:", resp)
	fmt.Println("body:", body)
}

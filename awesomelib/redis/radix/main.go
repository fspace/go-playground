package main

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
)

func main() {
	pool, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		// handle error
	}
	fmt.Println(pool)
}

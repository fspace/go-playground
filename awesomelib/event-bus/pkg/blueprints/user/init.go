package user

import (
	"fmt"
	"playgo/awesomelib/event-bus/pkg/core"
)

// BootStrap code  goes here  or we can name this file as bootstrap.go

func init()  {
	bus := core.Bus
	bus.Subscribe("main:loadconfig", func(data string) {
		fmt.Println("user:Subscribe(main:loadconfig)",data)
	})
}

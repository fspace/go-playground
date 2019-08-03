package main

import (
	events "github.com/kataras/go-events"
)

/**
分布式版本？
https://github.com/emitter-io/emitter
*/

func main() {
	// initialize a new EventEmmiter to use
	e := events.New()

	// register an event with name "my_event" and one listener
	e.On("my_event", func(payload ...interface{}) {
		message := payload[0].(string)
		print(message) // prints "this is my payload"
	})

	// fire the 'my_event' event
	e.Emit("my_event", "this is my payload")
}

package main

import (
	"fmt"
	"playgo/awesomelib/event-bus/pkg/core"

	// evbus "github.com/asaskevich/EventBus"
	_ "playgo/awesomelib/event-bus/pkg/blueprints/user"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a + b)
}
/**
## 相似库
- "github.com/sadlil/go-trigger"
 */
func main() {
	bus := core.Bus
	bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	bus.Unsubscribe("main:calculator", calculator)

	bus.Publish("main:loadconfig","hi this is the data from config file?")
}
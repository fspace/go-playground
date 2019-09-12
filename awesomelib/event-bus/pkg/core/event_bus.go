package core

import (
	// evbus "github.com/asaskevich/EventBus"
	"github.com/asaskevich/EventBus"
)

// 跨项目的可以考虑： https://github.com/ThreeDotsLabs/watermill
// 还有一个奇怪语法的： https://github.com/olebedev/emitter

func init() {
	Bus = EventBus.New()
}

var Bus EventBus.Bus

/**
func calculator(a int, b int) {
	fmt.Printf("%d\n", a + b)
}

## 相似库
- "github.com/sadlil/go-trigger"

func main() {
	bus := EventBus.New()
	bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	bus.Unsubscribe("main:calculator", calculator)
}
*/

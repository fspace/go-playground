package main

import (
	"github.com/mustafaturan/bus"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

func init() {
	// configure id generator (it doesn't have to be monoton)
	node := uint(1)
	initialTime := uint(0)
	monoton.Configure(sequencer.NewMillisecond(), node, initialTime)

	// configure bus
	if err := bus.Configure(bus.Config{Next: monoton.Next}); err != nil {
		panic("whoops")
	}
	// ...
	// register topics
	bus.RegisterTopics("order.received", "order.fulfilled")

	handler := bus.Handler{
		Handle: func(e *Event) {
			// fmt.Printf("Event: %+v %+v\n", e, e.Topic)
			// do something
			// NOTE: Highly recommended to process the event in an async way
		},
		Matcher: ".*", // regex pattern that matches all topics
	}
	bus.RegisterHandler("a unique key for the handler", &handler)
}

func main() {
	txID := "some-transaction-id-if-exists" // if it is blank, bus will generate one
	topic := "order.received"               // event topic name (must be registered before)
	order := make(map[string]string)        // interface{} data for event
	order["orderID"] = "123456"
	order["orderAmount"] = "112.20"
	order["currency"] = "USD"

	bus.Emit(topic, order, txID) // emit the event for the topic
}

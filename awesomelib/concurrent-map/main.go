package main

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
)

// 类似库： https://github.com/modern-go/concurrent

func main() {
	// Create a new map.
	m := cmap.New()

	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")

	// Retrieve item from map.
	if tmp, ok := m.Get("foo"); ok {
		bar := tmp.(string)

		fmt.Printf("bar type is %T , and value is : %v \n", bar, bar)
	}

	// Removes item under key "foo"
	m.Remove("foo")
}

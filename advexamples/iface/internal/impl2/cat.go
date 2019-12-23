package impl2

import (
	"fmt"
	"playgo/advexamples/iface/internal"
)

var _ internal.Duck = &Cat{}

type Cat struct {
}

func (c *Cat) Walk() {
	fmt.Println("ptr cat walk ")
}

func (c *Cat) Quack() {
	fmt.Println("ptr cat : meow")
}

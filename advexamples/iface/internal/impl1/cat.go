package impl1

import (
	"fmt"
	"playgo/advexamples/iface/internal"
)

var _ internal.Duck = Cat{}

type Cat struct {
}

func (Cat) Walk() {
	fmt.Println("cat walk")
}

func (Cat) Quack() {
	fmt.Println("meow!")
}

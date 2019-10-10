package main

import "fmt"

func main() {
	c, _ := New(Name("some-component"))
	c.Inspect()
}

// --------------------------------------------------------

// New returns a new Component.
func New(opts ...Option) (*MyComponent, error) {
	// The default configuration.
	conf := &MyConfig{
		Author: "yiqing",
	}
	for _, o := range opts {
		o(conf)
	}

	c := &MyComponent{
		Id:   conf.Name,
		Size: conf.Size,
	}
	return c, nil
}

type MyComponent struct {
	Id   string
	Size int
}

func (c *MyComponent) Clone(opts ...Option) *MyComponent {
	conf := &MyConfig{
		Size: c.Size,
	}
	for _, o := range opts {
		o(conf)
	}

	clone := &MyComponent{
		Size: conf.Size,
	}
	return clone
}

func (c *MyComponent) Inspect() {
	fmt.Printf("%#v", c)
}

type MyConfig struct {
	Name   string
	Author string
	Size   int
}
type Option func(config *MyConfig)

// Name
func Name(name string) Option {
	return Option(func(c *MyConfig) {
		c.Name = name
	})
}

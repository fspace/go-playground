package console

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Context for "ls" command
type AddUserCommand struct {
	//All bool
	Data string
}

func (cmd *AddUserCommand) Run(c *kingpin.ParseContext) error {
	//fmt.Printf("all=%v\n", l.All)
	fmt.Println("add user -->")
	//fmt.Printf("context is : %#v",c)
	// fmt.Printf("context is : %v",c.Elements)
	fmt.Printf("data : %v",cmd.Data)
	return nil
}

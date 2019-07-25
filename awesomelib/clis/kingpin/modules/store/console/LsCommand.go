package console

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Context for "ls" command
type LsCommand struct {
	All bool
}

func (l *LsCommand) Run(c *kingpin.ParseContext) error {
	fmt.Printf(" store: all=%v\n", l.All)
	return nil
}

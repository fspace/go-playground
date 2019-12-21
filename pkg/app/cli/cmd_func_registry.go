package cli

import (
	"fmt"
	"github.com/mitchellh/cli"
)

// copy from https://github.com/hashicorp/consul/blob/master/command/registry.go
// VoidFunc is a function that will be executed as a new instance of a CLI-sub command.
type VoidFunc func()

// Register adds a new CLI sub-command to the registry.
func Register(name string, fn VoidFunc) {
	if registry == nil {
		registry = make(map[string]VoidFunc)
	}

	if registry[name] != nil {
		panic(fmt.Errorf("func %q is already registered", name))
	}
	registry[name] = fn
}

// Map returns a realized mapping of available CLI commands in a format that
// the CLI class can consume. This should be called after all registration is
// complete.
func Map() map[string]cli.CommandFactory {
	m := make(map[string]cli.CommandFactory)
	for name, fn := range registry {
		thisFn := fn
		m[name] = func() (cli.Command, error) {
			// return thisFn(ui)
			return &DefaultCmd{
				RunFunc: thisFn,
			}, nil
		}
	}
	return m
}

// registry has an entry for each available CLI sub-command, indexed by sub
// command name. This should be populated at package init() time via Register().
var registry map[string]VoidFunc

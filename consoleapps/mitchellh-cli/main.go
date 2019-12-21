package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

// 改自MockCommand 也可以起名为RelaxCommand|FreeCommand
type BaseCmd struct {
	// Settable
	HelpText     string
	RunResult    int
	SynopsisText string
	//
	RunFunc   func()                  // 无参类型的函数类型 -- 真正执行的函数
	CustomRun func(args []string) int // 替换掉默认实现 TODO 暂未提供实现 替换 委托给原有实现即可

	// Set by the command
	RunCalled bool
	RunArgs   []string
}

func (c *BaseCmd) Help() string {
	return c.HelpText
}

func (c *BaseCmd) Run(args []string) (code int) {
	// 我们需要在recover分支流程中 修改返回值 所以用了命名返回值 https://blog.golang.org/defer-panic-and-recover
	if c.CustomRun != nil {
		return c.CustomRun(args)
	}
	c.RunCalled = true
	c.RunArgs = args

	// 执行真正的代码
	if c.RunFunc != nil {
		defer func() {
			// Deferred functions may read and assign to the returning function's named return values.
			if err := recover(); err != nil {
				// c.RunResult = 2
				code = 222
				c.RunResult = 222
				//fmt.Println("panic: lalalalalalla...",c.RunResult)
				fmt.Println("panic la ...", err)
			}
		}()
		c.RunFunc()
	}

	return c.RunResult
}

func (c *BaseCmd) Synopsis() string {
	return c.SynopsisText
}

func main() {
	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		//"foo": fooCommandFactory,
		//"bar": barCommandFactory,
		"foo": func() (cli.Command, error) {
			return &BaseCmd{
				RunFunc: func() {
					fmt.Println("hi  from real func!")
					panic("hello ya , panic from some fun! ")
				},
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	// fmt.Println("exit-status: ", exitStatus)
	os.Exit(exitStatus)
}

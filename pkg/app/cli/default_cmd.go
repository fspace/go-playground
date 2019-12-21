package cli

import "fmt"

// 改自MockCommand 也可以起名为RelaxCommand|FreeCommand
type DefaultCmd struct {
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

func (c *DefaultCmd) Help() string {
	return c.HelpText
}

func (c *DefaultCmd) Run(args []string) (code int) {
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

func (c *DefaultCmd) Synopsis() string {
	return c.SynopsisText
}

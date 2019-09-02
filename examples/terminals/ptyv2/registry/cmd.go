package registry

import (
	"fmt"
	"io"
	"playgo/examples/terminals/ptyv2/pkg"
)

// 注册表作为共享包 里面的注册项可用模拟 os注册表的概念 做分类

var cmds []pkg.Cmd

func init() {
	RegisterCmd(help)
}

func RegisterCmd(cmd pkg.Cmd) {
	cmds = append(cmds, cmd)
}

func GetCmds() []pkg.Cmd {
	return cmds
}

var help = pkg.Cmd{
	Name: "Help", // NOTE 大小写必须严格匹配么？
	Help: "Shows available cmds ",
	Action: func(w io.Writer, args ...string) bool {
		fmt.Fprintln(w, "Available commands:")
		for _, c := range cmds {
			fmt.Fprintf(w, " - %-15s %s\n", c.Name, c.Help)
		}

		return false
	},
}

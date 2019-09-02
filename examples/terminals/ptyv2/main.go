package main

import (
	"bufio"
	"fmt"
	"os"
	"playgo/examples/terminals/ptyv2/registry"
	"strings"
)

// @see Hands-On System Programming with Go

// pseudo-teletypes (PTY)
func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Fprint(w, "Some welcome message\n")
	for {
		s.Scan() // get next the token
		//msg := string(s.Bytes())
		//if msg == "exit" {
		//	return
		//}
		//fmt.Fprintf (w, "You wrote %q\n", msg) // writing back the text

		args := strings.Split(string(s.Bytes()), " ")
		cmd := args[0]
		args = args[1:]

		cmds := registry.GetCmds()
		fmt.Println("cmds ", cmds)

		var idx int = -1
		for i := range cmds {
			// 这里有责任链的影子！
			if !cmds[i].Match(cmd) {
				continue
			}
			idx = i
			break
		}

		if idx == -1 {
			//fmt.Fprintf(w, "%q not found. Use `help` for available commands\n", args[0])
			fmt.Fprintf(w, "%s not found. Use `help` for available commands\n", cmd)
			continue
		}
		//if cmds[idx].Run(w, args[1:]...) {
		if cmds[idx].Run(w, args...) {
			fmt.Fprintln(w)
			return
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"playgo/examples/terminals/pty/cmds"
	"strings"
)

// @see Hands-On System Programming with Go

var cmdFunc func(w io.Writer, args []string) (exit bool)

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
		switch cmd {
		case "exit":
			cmdFunc = exitCmd
		case "shuffle":
			cmdFunc = cmds.Shuffle
		case "print":
			cmdFunc = cmds.Print

		default:
			fmt.Println("we can do this \n case \"someCmd\": \n someCmd(w,args) \n ")
			fmt.Fprintf(w, " cmd:[%s]  \t args: %v \n ", cmd, args)

		}

		if cmdFunc == nil {
			fmt.Fprintf(w, "%q not found \n", cmd)
			continue
		}

		if cmdFunc(w, args) {
			// execute and exit if true
			return
		}
		// 清零  防止“记忆"
		cmdFunc = nil
	}
}

func exitCmd(w io.Writer, args []string) bool {
	fmt.Fprintf(w, "Goodby!: ")
	return true
}

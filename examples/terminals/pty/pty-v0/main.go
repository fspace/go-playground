package main

import (
	"bufio"
	"fmt"
	"os"
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
		switch cmd {
		case "exit":
			return
		default:
			fmt.Println("we can do this \n case \"someCmd\": \n someCmd(w,args) \n ")
			fmt.Fprintf(w, " cmd:[%s]  \t args: %v \n ", cmd, args)

		}

	}
}

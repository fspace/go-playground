package cmds

import (
	"fmt"
	"io"
	"os"
)

//func Print(w io.Writer, args ...string)bool   {
func Print(w io.Writer, args []string) bool {
	if len(args) != 1 {
		fmt.Fprintln(w, "Pls specify one file!")
		return false
	}

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(w, "Cannot open %s: %s \n", args[0], err)
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		fmt.Fprintf(w, "Cannot print %s: %s \n", args[0], err)
	}
	fmt.Println(w)
	return false
}

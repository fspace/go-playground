package cmds

import (
	"fmt"
	"io"
	"math/rand"
)

func Shuffle(w io.Writer, args []string) bool {
	fmt.Fprintf(w, "Shuffle cmd : args %v \n", args)

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
	})
	for i := range args {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%s", args[i])
	}
	fmt.Fprintln(w)
	return false
}

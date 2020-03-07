package main

/**
- https://fasterthanli.me/blog/2020/i-want-off-mr-golangs-wild-ride
*/
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arg := os.Args[1]
	fi, _ := os.Stat(arg)
	fmt.Printf("(%s) mode = %o\n", arg, fi.Mode()&os.ModePerm)
}

func main2() {
	arg := os.Args[1]

	fi, err := os.Stat(arg)
	must(err)
	fmt.Printf("(%s) old mode = %o\n", arg, fi.Mode()&os.ModePerm)

	must(os.Chmod(arg, 0755))

	fi, err = os.Stat(arg)
	must(err)
	fmt.Printf("(%s) new mode = %o\n", arg, fi.Mode()&os.ModePerm)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// -------------------------------------------------------------------------------------

func main3() {
	arg := os.Args[1]
	f, err := os.Open(arg)
	must(err)

	entries, err := f.Readdir(-1)
	must(err)

	for _, e := range entries {
		if e.IsDir() {
			fmt.Printf("(dir) %s\n", e.Name())
		} else {
			fmt.Printf("      %s\n", e.Name())
		}
	}
}

// -------------------------------------------------------------------------------------
func main4() {
	inputs := []string{
		"/",
		"/.",
		"/.foo",
		"/foo",
		"/foo.txt",
		"/foo.txt/bar",
		"C:\\",
		"C:\\.",
		"C:\\foo.txt",
		"C:\\foo.txt\\bar",
	}
	for _, i := range inputs {
		fmt.Printf("%24q => %q\n", i, filepath.Ext(i))
	}
}

// -------------------------------------------------------------------------------------

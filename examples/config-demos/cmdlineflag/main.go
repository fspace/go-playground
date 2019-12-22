//cmdlineflag.go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	// main operation modes
	write = flag.Bool("w", false, "write result back instead of stdout\n\t\tDefault: No write back")

	// layout control
	tabWidth = flag.Int("tabwidth", 8, "tab width\n\t\tDefault: Standard")

	// debugging
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file\n\t\tDefault: no default")
)

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flags] file [path ...]\n\n",
		"CommandLineFlag") // os.Args[0]
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	fmt.Printf("Before parsing the flags\n")
	fmt.Printf("T: %d\nW: %s\nC: '%s'\n",
		*tabWidth, strconv.FormatBool(*write), *cpuprofile)

	flag.Usage = usage
	flag.Parse()

	// There is also a mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		usage()
	}

	fmt.Printf("Testing the flag package\n")
	fmt.Printf("T: %d\nW: %s\nC: '%s'\n",
		*tabWidth, strconv.FormatBool(*write), *cpuprofile)

	for index, element := range flag.Args() {
		fmt.Printf("I: %d C: '%s'\n", index, element)
	}
}

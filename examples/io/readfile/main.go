package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please specify a file")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer f.Close()

	var (
		b = make([]byte, 16)
	)
	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Print(string(b[:n])) // only print what's been read
		}
	}
	if err != nil && err != io.EOF {
		// 我们期待的是 EOF 错误
		fmt.Println("\n\n Error: ", err)
	}

}

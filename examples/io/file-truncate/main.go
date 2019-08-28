package main

import (
	"fmt"
	"os"
)

func main() {
	// 低于 4kb
	if err := os.Truncate("file.txt", 4096); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

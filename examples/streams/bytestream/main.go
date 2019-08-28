package main

import "fmt"

func main() {
	const a = `⌘` // 这是一个字符串！
	fmt.Printf("plain str: %s\n", a)
	fmt.Printf("quoted str: %q\n", a)
	fmt.Println("hex bytes:")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%x", a[i])
	}
	fmt.Printf("\n")
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "A", "Simple", "cmd")
	fmt.Println(cmd.Path, cmd.Args[1:])

	demo1()
	// demo2()
}

func demo1() {
	// 奇怪的错误 怎么没有dir呢？
	// cmd := exec.Command("dir")
	cmd := exec.Command("ls", "-l") // win 下面没有
	if err := cmd.Start(); err != nil {
		// 异步执行
		log.Fatalln(err)
	}
	fmt.Println("Cmd: ", cmd.Args[0])
	fmt.Println("Cmd: ", cmd.Args[1:])
	fmt.Println("PID: ", cmd.Process.Pid)

	cmd.Wait()
}

func demo2() {
	b := bytes.NewBuffer(nil)
	cmd := exec.Command("cat")
	cmd.Stdin = b
	cmd.Stdout = os.Stdout
	fmt.Fprintf(b, "Hello World! I'm using this memory address: %p", b)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	cmd.Wait()
}

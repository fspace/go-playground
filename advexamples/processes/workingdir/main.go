package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Working dir: ", wd)
	fmt.Println("Application: ", filepath.Join(wd, os.Args[0]))
	// create a new dir
	d := filepath.Join(wd, "test")
	// fmt.Println(d)
	if err := os.Mkdir(d, 0755); err != nil {
		fmt.Println("Err: ", err)
		return
	}
	fmt.Println("Created : ", d)
	// change the current directory
	if err := os.Chdir(d); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Working Directory: ", d)

}

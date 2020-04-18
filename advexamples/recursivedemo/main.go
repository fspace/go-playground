package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("recursive-demo", "demo for recursive data structures and functions  ")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
type storyPage struct {
	text     string
	nextPage *storyPage
}

// playStory is a recursion function which play the given page one after another !
func playStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

// ==============================================================================
func basics() {

	p1 := storyPage{"this is page1: It was a dark and stormy night", nil}
	p2 := storyPage{"You are alone , and you need find the sacred helmet before the bad guys do", nil}
	p3 := storyPage{"You see a troll ahead ", nil}
	p1.nextPage = &p2
	p2.nextPage = &p3

	playStory(&p1)
}

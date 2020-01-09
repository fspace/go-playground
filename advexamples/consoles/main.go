package main

import (
	"bufio"
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"strings"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("r", "Read Rune", cli.ActionCommand(ReadRune))
	app.Command("rl", "Read new line ", cli.ActionCommand(Readline))
	app.Command("rc", "Read In From Console", cli.ActionCommand(ReadInFromConsole))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
/**
If you are only needing single character input then use ReadRune() or if you are wanting to read in full new line
delimited sentences then ReadString is the way to go
*/
// ==============================================================================
func ReadRune() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}
}

func Readline() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		// text = strings.Replace(text, "\n", "", -1)
		// win OS
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		} else {
			fmt.Println("Your input is : ", text)
		}

	}

}

func ReadInFromConsole() {
	/**
	infinitely ask scan for input and echo back whatever is entered.
	*/
	scanner()
}
func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

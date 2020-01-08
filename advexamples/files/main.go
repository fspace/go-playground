package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点

	app.Command("r", "Reading Files", cli.ActionCommand(ReadingFiles))
	app.Command("w", "Writing Files", cli.ActionCommand(WritingFiles))
	app.Command("a", "Appending Data", cli.ActionCommand(AppendingData))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================
func ReadingFiles() {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("localfile.data")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Print(string(data))
}

func WritingFiles() {
	mydata := []byte("all my data I want to write to a file" + strconv.FormatInt(time.Now().UnixNano(), 10))
	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.data", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
	fmt.Println("ok!")

	// 再读出来
	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}

func AppendingData() {
	f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}

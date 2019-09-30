package main

import (
	"log"
	"os"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
			errorLog.Fatalln("no: ", e)
		}
	}()
	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	// errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("main running ...")

	panic("some panic")
	infoLog.Println("cannot goes here ")
}

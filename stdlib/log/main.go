package main

import (
	"log"
	"os"
)

/**
As a rule of thumb, you should avoid using the Panic() and Fatal() variations outside of your main() function —
it’s good practice to return errors instead, and only panic or exit directly from main().
- log.New 生出来的日志对象是并发安全的
*/

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
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // 可以追踪源码看下还有其他可用的flag哦！ Llongfile

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	// errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("main running ...")

	// 测试下日志打到文件 虽然推荐做法是在运行时用标准输出做重定向到其他流 这里也不妨试试直接输出到文件：
	playLog()

	panic("some panic")
	infoLog.Println("cannot goes here ")
}

func playLog() {
	f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Println("some info string  , log to file: ", f.Name())
}

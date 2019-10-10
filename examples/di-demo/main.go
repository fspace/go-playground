package main

import (
	"log"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	_ = app
	// 其他包会依赖不同的全局组件比如： logger ，cache dbConnection 等 把他们收集到Application结构中 然后注入给他们
	// 对于依赖 要么是自己拉取 比如 搞个全局注册表包 register|config|g.        要么是注入到不同的包 xxx.Init(app)
	/**
	    // 闭包式用法
		app := &config.Application{
			ErrorLog:log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)
		}

		mux.Handle("/",handlers.Home(app))
	*/
}

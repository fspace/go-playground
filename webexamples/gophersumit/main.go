package main

import (
	"github.com/jawher/mow.cli"
	"net/http"
	"os"
	"playgo/webexamples/gophersumit/internal/handlers"
	"playgo/webexamples/gophersumit/internal/multiplexer"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("ss", "using server struct", cli.ActionCommand(usingServerStruct))
	app.Command("m", "multiplexer", cli.ActionCommand(multiplexer.Main))
	//handlers in Go using two ways
	app.Command("st", "struct type as handler", cli.ActionCommand(structTypeAsHandler))
	app.Command("fah", "function as handler", cli.ActionCommand(functionAsHandler))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gophers!"))
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":3000", mux)
}

func usingServerStruct() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello,2  Gophers!"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	httpServer.ListenAndServe()
}

func structTypeAsHandler() {
	hand := handlers.CustomHandler{}
	mux := http.NewServeMux()
	mux.Handle("/", &hand)

	http.ListenAndServe(":3000", mux)
}

func functionAsHandler() {
	functionHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("function as http handler!"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", functionHandler)
	http.ListenAndServe(":3000", mux)
}

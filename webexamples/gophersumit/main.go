package main

import (
	"github.com/jawher/mow.cli"
	"net/http"
	"os"
	"playgo/webexamples/gophersumit/internal/handlers"
	"playgo/webexamples/gophersumit/internal/multiplexer"
	"playgo/webexamples/gophersumit/internal/part2/headers"
	"playgo/webexamples/gophersumit/internal/part2/query"
	"playgo/webexamples/gophersumit/internal/part2/request"
	"playgo/webexamples/gophersumit/internal/part2/response"
	"playgo/webexamples/gophersumit/internal/part3/cookies"
	"playgo/webexamples/gophersumit/internal/part3/sessions"
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

	app.Command("part2/request", "extract multiple details from incoming request ", cli.ActionCommand(request.Main))
	app.Command("part2/response", "build the http response using ResponseWriter interface. ", cli.ActionCommand(response.Main))
	app.Command("part2/header", "see headers coming from incoming http request. ", cli.ActionCommand(headers.Main))
	app.Command("part2/query", " get query string values ", cli.ActionCommand(query.QueryStrings))

	app.Command("part3/set-cookie", "  set cookies ", cli.ActionCommand(cookies.Main))
	app.Command("part3/get-cookie", "  get cookies ", cli.ActionCommand(cookies.GetCookies))
	app.Command("part3/session", "  get and set session data ", cli.ActionCommand(sessions.Main))
	app.Command("part3/flash", "  work with flash messages ", cli.ActionCommand(sessions.FlashMessage))

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

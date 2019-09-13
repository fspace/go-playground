package app

import (
	"net/http"
)

// @see https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

type server struct {
	//db     *someDatabase
	//router *someRouter
	//email  EmailSender

	// Shared dependencies are fields of the structure
}

func (s *server) routes() {
	//s.router.HandleFunc("/api/", s.handleAPI())
	//s.router.HandleFunc("/about", s.handleAbout())
	//s.router.HandleFunc("/", s.handleIndex())

	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}

func (s *server) handleSomething() http.HandlerFunc {
	//thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

// Take arguments for handler-specific dependencies
//func (s *server) handleGreeting(format string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, format, "World")
//	}
//}

// Middleware are just Go functions
//func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if !currentUser(r).IsAdmin {
//			http.NotFound(w, r)
//			return
//		}
//		h(w, r)
//	}
//}

// ## Request and response types can go in there too
//func (s *server) handleSomething() http.HandlerFunc {
//	type request struct {
//		Name string
//	}
//	type response struct {
//		Greeting string `json:"greeting"`
//	}
//	return func(w http.ResponseWriter, r *http.Request) {
//		...
//	}
//}

// ## sync.Once to setup dependencies
//
//func (s *server) handleTemplate(files string...) http.HandlerFunc {
//	var (
//		init sync.Once
//		tpl  *template.Template
//		err  error
//	)
//	return func(w http.ResponseWriter, r *http.Request) {
//		init.Do(func(){
//			tpl, err = template.ParseFiles(files...)
//		})
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		// use tpl
//	}
//}

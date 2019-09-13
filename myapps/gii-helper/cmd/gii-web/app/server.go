package app

import (
	"net/http"
)

// @see https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

type Server struct {
	AppConfig *Config
	//db     *someDatabase
	//router *someRouter
	//email  EmailSender

	// Shared dependencies are fields of the structure
	Router *http.ServeMux
}

func (s *Server) Routes() {
	//s.router.HandleFunc("/api/", s.handleAPI())
	//s.router.HandleFunc("/about", s.handleAbout())
	//s.router.HandleFunc("/", s.handleIndex())

	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
	s.Router.HandleFunc("/", s.handleIndex())
}

func (s *Server) handleIndex() http.HandlerFunc {
	//thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
		w.Header().Set("Content-Type", "text/html")
		//     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		html := `
		<p>hello </p>
`
		w.Write([]byte(html))
	}
}

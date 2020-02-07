package multiplexer

import "net/http"

var mux = http.NewServeMux()

func registerRoutes() {
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
}

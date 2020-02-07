package multiplexer

import "net/http"

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about route"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home route"))
}
func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login route"))
}
func logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout route"))
}

package ch6

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Main() {
	router := NewRouter()

	router.Handle("GET", "/", HandleHome)
	router.Handle("GET", "/register", HandleUserNew)
	router.Handle("POST", "/register", HandleUserCreate)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("internal/ch6/assets/"),
	)

	middleware := Middleware{}
	middleware.Add(router)
	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	return router
}

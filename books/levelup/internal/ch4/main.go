package ch4

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	//"github.com/julienschmidt/httprouter"
)

func Main() {
	unauthenticatedRouter := NewRouter()
	unauthenticatedRouter.GET("/", HandleHome)
	authenticatedRouter := NewRouter()
	authenticatedRouter.GET("/images/new", HandleImageNew)
	middleware := Middleware{}
	middleware.Add(unauthenticatedRouter)
	middleware.Add(http.HandlerFunc(AuthenticateRequest))
	middleware.Add(authenticatedRouter)
	http.ListenAndServe(":3000", middleware)
}

// Creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	return router
}

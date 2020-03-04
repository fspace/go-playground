package ch4

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleUserNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Display Home Page
	RenderTemplate(w, r, "users/new", nil)
}

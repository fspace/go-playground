package ch6

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Display Home Page
	RenderTemplate(w, r, "index/home", nil)
}

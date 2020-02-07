package handlers

import "net/http"

type CustomHandler struct{}

// ServeHTTP            satisfy Handler interface
func (hand *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom handler!"))
}

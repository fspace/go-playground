package main

import (
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      mux,
	}
	server.ListenAndServe()
}

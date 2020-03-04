package ch5

import (
	"log"
	"net/http"
)

func SimpleStaticWebserver() {
	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("internal/ch5/assets/"))))
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func SimpleServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// we can conditional check post or get request
		RenderTemplate(w, r, "index/home", nil)
	})
	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("internal/ch5/assets/"))),
	)
	http.ListenAndServe(":3000", mux)
}

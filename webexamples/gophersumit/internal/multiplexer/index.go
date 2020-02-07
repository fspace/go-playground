package multiplexer

import "net/http"

func Main() {
	registerRoutes()
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}

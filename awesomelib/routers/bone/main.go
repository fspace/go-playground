package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", bone.GetValue(r, "id"))
}

func main() {
	mux := bone.New()
	mux.Get("/some/page/:id", http.HandlerFunc(MyHandler))
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

package headers

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	hed := r.Header // type Header map[string][]string

	fmt.Fprintln(w, hed)
	// 两种获取header单个值的方法：
	hed1 := r.Header.Get("Accept")
	fmt.Fprintln(w, hed1)

	hed2 := r.Header["Accept"]
	fmt.Fprintln(w, hed2)
}

func setHeader(w http.ResponseWriter, r *http.Request) {
	// NOTE: Unlike name suggests, WriteHeader() on ResponseWriter is not used to set headers in response. net/http package has documented this part:
	// WriteHeader sends an HTTP response header with the provided
	// status code.
	w.Header().Set("ALLOWED", "GET, POST")
	w.Write([]byte("set allowed headers\n"))
}

func setHeader2(w http.ResponseWriter, r *http.Request) {
	//  send http status code other than 200
	// NOTE:It is important to call WriteHeader before Write if status code we want to send is other than 200.
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request!\n"))
}

func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", headers)
	mux.HandleFunc("/set-header", setHeader)
	mux.HandleFunc("/write-header", setHeader2)
	http.ListenAndServe(":3000", mux)
}

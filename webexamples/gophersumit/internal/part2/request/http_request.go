package request

import (
	"net/http"
	"strconv"
)

func requestInspection(w http.ResponseWriter, r *http.Request) {
	/**
	- Body of request
	- Query string parameters
	- HTTP Headers
	- Posted form etc.
	*/
	w.Write([]byte(string("Method: " + r.Method + "\n")))
	w.Write([]byte(string("Protocol Version: " + r.Proto + "\n")))
	w.Write([]byte(string("Host: " + r.Host + "\n")))
	w.Write([]byte(string("Referer: " + r.Referer() + "\n")))
	w.Write([]byte(string("User Agent: " + r.UserAgent() + "\n")))
	w.Write([]byte(string("Remote Addr: " + r.RemoteAddr + "\n")))
	w.Write([]byte(string("Requested URL: " + r.RequestURI + "\n")))
	w.Write([]byte(string("Content Length: " + strconv.FormatInt(r.ContentLength, 10) + "\n")))
	for key, value := range r.URL.Query() {
		w.Write([]byte(string("Query string: key=" + key + " value=" + value[0] + "\n")))
	}

}

func Main() {
	// curl http://localhost:3000/home?name=sumit
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestInspection)

	http.ListenAndServe(":3000", mux)
}

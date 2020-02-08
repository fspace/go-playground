package query

import "net/http"

//
func showQuery(w http.ResponseWriter, r *http.Request) {
	/**
		type Request struct {
		  // other field omitted
		  URL *url.URL
		}
	  // -----------------
		type URL struct {
		  // fields omitted
		  RawQuery   string    // encoded query values, without '?'
		}

		// Query parses RawQuery and returns the corresponding values.
		// It silently discards malformed value pairs.
		// To check errors use ParseQuery.
		func (u *URL) Query() Values {
			v, _ := ParseQuery(u.RawQuery)
			return v
		}

	*/
	querystring := r.URL.Query()
	w.Write([]byte("query strings are\n"))
	w.Write([]byte("Name:" + querystring.Get("name") + "\n"))
	w.Write([]byte("Email:" + querystring.Get("email") + "\n"))
}
func QueryStrings() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", showQuery)
	http.ListenAndServe(":3000", mux)
}

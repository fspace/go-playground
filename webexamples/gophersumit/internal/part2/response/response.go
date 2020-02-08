package response

import "net/http"

func unauthorized(w http.ResponseWriter, r *http.Request) {
	// NOTE: When returning status code other than 200 we must call w.WriteHeader() before any call to w.Write().
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("you do not have permission to access this resource.\n"))
}
func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", unauthorized)

	http.ListenAndServe(":3000", mux)
}

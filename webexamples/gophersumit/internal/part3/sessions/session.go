package sessions

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// for prod use secure key, not hard-coded
var store = sessions.NewCookieStore([]byte("1234"))

func sessionHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "custom-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val := session.Values["hello"]
	if val != nil {
		fmt.Fprintln(w, "retrieving existing session: ")
		fmt.Fprintln(w, val)
		return
	}
	session.Values["hello"] = "world"
	session.Save(r, w)
	fmt.Fprintln(w, "no existing session found, set value for session")

}

func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sessionHandler)
	http.ListenAndServe(":3000", mux)
}

func FlashMessage() {
	// 什么原因 没有获取到flash  ？？？
	flashAdderHandler := func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "custom-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		message := "some op notifier message"
		session.AddFlash(message)
	}
	flashGetHandler := func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "custom-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		flashes := session.Flashes()
		fmt.Fprintln(w, flashes)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", flashAdderHandler)
	mux.HandleFunc("/get", flashGetHandler)
	http.ListenAndServe(":3000", mux)
}

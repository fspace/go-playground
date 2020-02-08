package cookies

import (
	"fmt"
	"net/http"
)

func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "cookie-1",
		Value: "hello world",
	}
	http.SetCookie(w, &cookie)
}

func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setCookies)
	http.ListenAndServe(":3000", mux)
}

// ............

func getCookies(w http.ResponseWriter, r *http.Request) {
	// get all cookies
	cookies := r.Cookies()
	for _, cookie := range cookies {
		fmt.Fprintln(w, cookie)
	}
	// get named cookie
	cookie, err := r.Cookie("cookie-1")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, cookie)
}

func GetCookies() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getCookies)
	http.ListenAndServe(":3000", mux)
}

/**
Since cookies are stored in user browser, they are generally used to enhance user experience.
*/

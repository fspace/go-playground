package basic

import (
	"fmt"
	"log"
	"net/http"
)

func Main() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Blog:www.flysnow.org\nwechat:flysnow_org")
}

package ch3

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// ======================================================================================
// UptimeHandler writes the number of seconds since starting the response
type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}
func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Current Uptime: %s", time.Since(h.Started)),
	)
}

// ---------------------------------------------------------------
// SecretTokenHandler secures a request with a secret token.
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

// ServeHTTP makes SecretTokenHandler implement the http.Handler interface.
func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Check the query string for the secret token
	if req.URL.Query().Get("secret_token") == h.secret {
		// The secret token matched, call the next handler
		h.next.ServeHTTP(w, req)
	} else {
		// No match, return a 404 Not Found response
		http.NotFound(w, req)
	}
}

type Article struct {
	Name       string
	AuthorName string
}

func (a Article) Byline() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

// ======================================================================================

func SimpleServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gopher")
	})
	http.ListenAndServe(":3000", nil)
}
func HttpError() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something has gone wrong", 500)
	})
	http.ListenAndServe(":3000", nil)
}
func HttpStatusHelpers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Return a 404 Not Found
		http.NotFound(w, r)
		// Return a 301 Permanently Moved
		// http.Redirect(w, req, "http://anothersite.com", 301)
	})
	http.ListenAndServe(":3000", nil)
}
func SimpleServer2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go Server")
		fmt.Fprintf(w, `<html>
<body>
Hello Gopher
</body>
</html>`)
	})
	http.ListenAndServe(":3000", nil)
}

func PathAndSubtrees() {
	/**
		A path is
	defined without a trailing backslash (/), and refers to an explicit path. Subtrees are
	designed to match the start of a path, and include the trailing /

		###
		The length of a pattern is important too. The longer a pattern is, the higher a precedence
	it has. A pattern of /articles/latest/ has a higher precedence than /articles/,
		和添加顺序无关！！！
	*/
	//mux := http.ServeMux{}
	//mux.Handle("/articles/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello from /articles/")
	//})
	//mux.Handle("/users", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello from /users")
	//})
	//http.ListenAndServe(":3000", mux)

	mux := http.NewServeMux()
	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/")
	})
	mux.HandleFunc("/articles/latest/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/latest/")
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /users")
	})
	http.ListenAndServe(":3000", mux)
}

func HandlerDemo() {
	http.Handle("/", NewUptimeHandler())
	http.ListenAndServe(":3000", nil)
}

func MiddlewareDemo() {
	// curl localhost:3000?secret_token=MySecret -i
	// HTTP/1.1 200 OK
	http.Handle("/", SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "MySecret",
	})
	http.ListenAndServe(":3000", nil)
}

func HtmlTemplates() {
	// Variable replacements, and control structures, are
	//referred to as actions and are surrounded by “{{” and “}}” characters

	tmpl, err := template.New("Foo").Parse("<h1>Hello {{.}}</h1>\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, "World")
	if err != nil {
		panic(err)
	}
}

func AccessingData() {
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}
	tmpl, err := template.New("Foo").Parse("'{{.Name}}' by " +
		"{{.AuthorName}}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}
}
func AccessingData2() {
	article := map[string]string{
		"Name":       "The Go html/template package",
		"AuthorName": "Mal Curtis",
	}
	tmpl, err := template.New("Foo").Parse("'{{.Name}}' by" +
		"➥{{.AuthorName}}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, article)
	if err != nil {
		panic(err)
	}
}

func NiladicFunction() {
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}
	tmpl, err := template.New("Foo").Parse("{{.Byline}}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

}

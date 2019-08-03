package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	toConsole()
	//toWeb_v2()
	//toWeb_v3()
	toWeb_v4()
}
func toWeb() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.
			New("go-web").
			//Parse(`hello world! {{.}}`)
			Parse(`
	Package : {{.Name}}
	Number of funcs : {{.NumFuncs}}
	Number of variables : {{.NumVars}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
		}
		err = tmpl.
			//Execute(os.Stdout,nil)
			//Execute(os.Stdout,"go web")
			Execute(w, &Package{
				Name:     "my package",
				NumFuncs: 12,
				NumVars:  12000,
			})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
		}
	})

	log.Println("Starting server at ：4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
func toWeb_v2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.
			ParseFiles("tpls/go-web.html")
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
		}
		err = tmpl.
			//Execute(os.Stdout,nil)
			//Execute(os.Stdout,"go web")
			Execute(w, &Package{
				Name:     "my package",
				NumFuncs: 12,
				NumVars:  12000,
			})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
		}
	})

	log.Println("Starting server at ：4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
func toWeb_v3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.
			ParseFiles("tpls/go-web.v3.html")
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
		}
		err = tmpl.
			//Execute(os.Stdout,nil)
			//Execute(os.Stdout,"go web")
			Execute(w, r)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
		}
	})

	log.Println("Starting server at ：4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
func toWeb_v4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.
			ParseFiles("tpls/go-web.v4.html")
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
		}
		err = tmpl.
			Execute(w, map[string]interface{}{
				"Request": r,
				"SomeKey": "some-value",
				"SomeInt": 10,
			})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
		}
	})

	log.Println("Starting server at ：4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func toConsole() {
	tmpl, err := template.
		New("go-web").
		//Parse(`hello world! {{.}}`)
		Parse(`
	Package : {{.Name}}
	Number of funcs : {{.NumFuncs}}
	Number of variables : {{.NumVars}}
`)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	err = tmpl.
		//Execute(os.Stdout,nil)
		//Execute(os.Stdout,"go web")
		Execute(os.Stdout, &Package{
			Name:     "my package",
			NumFuncs: 12,
			NumVars:  12000,
		})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

package main

import (
	"github.com/Masterminds/sprig"
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
{{ "hello!" | upper | repeat 5 }}

`
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Funcs(sprig.TxtFuncMap()).Parse(letter))

	// Execute the template
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		log.Println("executing template:", err)
	}

}

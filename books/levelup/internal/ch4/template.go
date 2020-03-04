package ch4

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield called inappropriately")
	},
}
var layout = template.Must(
	template.
		New("layout.html").
		Funcs(layoutFuncs).
		ParseFiles("internal/ch5/templates/layout.html"),
)
var templates = template.Must(template.New("t").ParseGlob("internal/ch4/templates/**/*.html"))

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}
	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate, name, err),
			http.StatusInternalServerError,
		)
	}
}

var errorTemplate = `
<html>
<body>
<h1>Error rendering template %s</h1>
<p>%s</p>
</body>
</html>
`

/**
We’re creating a new template instance
called t, and telling it to parse the files that match the glob templates/** / *.html
*/

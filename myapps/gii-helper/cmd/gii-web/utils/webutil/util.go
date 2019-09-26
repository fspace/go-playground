package webutil

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

// https://blog.csdn.net/String12/article/details/79503566

const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

//commonly used mime-types
const (
	applicationJson = "application/json"
	applicationXml  = "application/xml"
	textXml         = "text/xml"
)

func ServeJson(w http.ResponseWriter, v interface{}) {
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", applicationJson)
	w.Write(content)
}

// ReadJson will parses the JSON-encoded data in the http
// Request object and stores the result in the value
// pointed to by v.
func ReadJson(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

// ServeXml replies to the request with an XML
// representation of resource v.
func ServeXml(w http.ResponseWriter, v interface{}) {
	content, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write(content)
}

// ReadXml will parses the XML-encoded data in the http
// Request object and stores the result in the value
// pointed to by v.
func ReadXml(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return err
	}
	return xml.Unmarshal(body, v)
}

// ServeFormatted replies to the request with
// a formatted representation of resource v, in the
// format requested by the client specified in the
// Accept header.
func ServeFormatted(w http.ResponseWriter, r *http.Request, v interface{}) {
	accept := r.Header.Get("Accept")
	switch accept {
	case applicationJson:
		ServeJson(w, v)
	case applicationXml, textXml:
		ServeXml(w, v)
	default:
		ServeJson(w, v)
	}

	return
}

func DisplayWithFuncs(w http.ResponseWriter, funcs template.FuncMap, d map[string]interface{}, tpls ...string) {
	if len(tpls) == 0 {
		return
	}
	name := filepath.Base(tpls[0])
	t := template.Must(template.New(name).Funcs(funcs).ParseFiles(tpls...))
	t.ExecuteTemplate(w, name, d)
}

package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"playgo/myapps/gii-helper/cmd/gii-web/utils/webutil"
	"playgo/myapps/gii-helper/pkg"
	"strconv"

	//"playgo/myapps/gii-helper/pkg"
	"strings"
	"time"
)

// @see https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

type Server struct {
	AppConfig *Config
	//db     *someDatabase
	//router *someRouter
	//email  EmailSender

	// Shared dependencies are fields of the structure
	// Router *http.ServeMux
	Router *mux.Router
}

func (s *Server) Routes() {
	//s.router.HandleFunc("/api/", s.handleAPI())
	//s.router.HandleFunc("/about", s.handleAbout())
	//s.router.HandleFunc("/", s.handleIndex())
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
	// /articles/{category}/{id:[0-9]+}

	th := &timeHandler{format: time.RFC1123}
	s.Router.Handle("/time", th)

	s.Router.HandleFunc("/db/{dbName}/table/{table}", s.handleColumns())
	s.Router.HandleFunc("/", s.handleIndex())
	// s.Router.HandleFunc("/*", s.handleAny()) // catch all 好像不管用
	s.Router.NotFoundHandler = s.handleAny()
}

func (s *Server) handleIndex() http.HandlerFunc {
	//thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
		q := r.URL.Query()
		fmt.Printf("query: %#v \n", q)
		fmt.Printf("path: %#v \n", r.URL.RawPath)

		w.Header().Set("Content-Type", "text/html")
		//     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		html := `
		<p>hello </p>
`
		w.Write([]byte(html))
	}
}

func (s *Server) handleAny() http.HandlerFunc {
	//thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
		q := r.URL.Query()
		fmt.Printf("query: %#v \n", q)
		fmt.Printf("path: %#v \n", r.URL.Path)

		w.Header().Set("Content-Type", "text/html")
		//     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		html := `
		<p>hello go</p>
`
		w.Write([]byte(html))
	}
}

func (s *Server) handleColumns() http.HandlerFunc {
	//thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {

		// use thing
		vars := mux.Vars(r)
		fmt.Printf("vars: %#v \n", vars)
		//
		table := vars["table"]
		dbName := vars["dbName"]

		replaceFn := func(s, from, to string) string {
			if strings.Contains(s, from) {
				return strings.Replace(s, from, to, 1)
			}
			return s
		}
		// _ = replaceFn
		dbDriver := s.AppConfig.DbDriver
		dsFmt := s.AppConfig.DataSourceFmt
		var ds string
		ds = replaceFn(dsFmt, "<dbuser>", s.AppConfig.DbUser)
		ds = replaceFn(ds, "<dbpass>", s.AppConfig.DbPass)
		ds = replaceFn(ds, "<dbhost>", s.AppConfig.DbHost)
		ds = replaceFn(ds, "<dbport>", strconv.Itoa(s.AppConfig.DbPort))
		ds = replaceFn(ds, "<dbname>", dbName)

		itr := pkg.NewDBInteractor(pkg.DBOption{
			DriverName: dbDriver,
			// "root:@/test?charset=utf8"
			DSName: ds,
		})
		rslt, err := itr.GetColumnsForTable(table)
		if err != nil {
			webutil.ServeJson(w, struct {
				Error bool
				Msg   string
			}{
				Error: true,
				Msg:   err.Error(),
			})
			return
		}

		webutil.ServeJson(w, rslt)
	}
}

// ===================================================================================
// 					## 时间服务
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

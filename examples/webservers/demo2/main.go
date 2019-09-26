package main

import (
	"log"
	"net/http"
	"time"
)

// @see https://blog.csdn.net/lengyuezuixue/article/details/79125541
// @see https://www.yuque.com/docs/share/c03cf792-227a-4e7a-850b-4789e479361d

// ## 中间件写法
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Comleted %s in %v", r.URL.Path, time.Since(start))
	})
}

// 环绕式
func hook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("before hook")
		next.ServeHTTP(w, r)
		log.Println("after hook")

	})
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      mux,
	}
	server.ListenAndServe()
}

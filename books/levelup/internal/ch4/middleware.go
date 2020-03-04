package ch4

import "net/http"

type Middleware []http.Handler

// Adds a handler to the middleware
func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Process the middleware
	mw := NewMiddlewareResponseWriter(w)
	// Loop through all the registered handlers
	for _, handler := range m {
		// Call the handler with our MiddlewareResponseWriter
		handler.ServeHTTP(mw, r)
		// If there was a write, stop processing
		if mw.written {
			return
		}
	}
	// If no handlers wrote to the response, it’s a 404
	http.NotFound(w, r)
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter // 内嵌一个接口类型 那么本类型就实现了这个接口哦
	written             bool
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter{
		ResponseWriter: w,
	}
}

// Write 复写内嵌实现的同名方法 但提供标记功能
func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

// WriteHeader 当写头时 标记写行为
func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

package http

import (
	"net/http"
)

type Router struct {
	mux        *http.ServeMux
	middleware Middleware
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Use(m Middleware) {
	r.middleware = m
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
	wrappedHandler := handler
	if r.middleware != nil {
		wrappedHandler = r.middleware(handler)
	}
	r.mux.HandleFunc(pattern, wrappedHandler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

package http

import (
	"fmt"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s '%s'\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

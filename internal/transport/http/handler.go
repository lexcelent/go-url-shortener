package http

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy"}`)
}

func DummyHandler(w http.ResponseWriter, req *http.Request) {}

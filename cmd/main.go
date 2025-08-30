package main

import (
	"log"
	"net/http"

	httprest "github.com/lexcelent/go-url-shortener/internal/transport/http"
)

func main() {
	router := httprest.NewRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Ошибка запуска HTTP-сервера")
	}
}

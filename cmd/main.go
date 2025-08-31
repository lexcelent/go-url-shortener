package main

import (
	"fmt"
	"log"
	"net/http"

	httprest "github.com/lexcelent/go-url-shortener/internal/transport/http"
)

func main() {
	router := httprest.NewRouter()
	router.Use(httprest.LoggingMiddleware)
	router.Handle("/api/health", httprest.HealthHandler)
	router.Handle("/dummy", httprest.DummyHandler)
	router.Handle("/url", httprest.UrlShort)

	fmt.Printf("Сервер запущен\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Ошибка запуска HTTP-сервера\n")
	}
}

package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lexcelent/go-url-shortener/internal/model"
	"github.com/lexcelent/go-url-shortener/internal/services"
)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy"}`)
}

func DummyHandler(w http.ResponseWriter, req *http.Request) {}

func UrlShort(w http.ResponseWriter, req *http.Request) {
	// Валидация запроса
	if req.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг тела запроса
	var msg model.Msg
	err := json.NewDecoder(req.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	newUrl := services.Short()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": newUrl,
	})
}

// TODO: Handle redirect
// TODO: redirect logic

// TODO: move structs to internal/domain/model
// TODO: make ENV

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

	// Работа с бизнес логикой
	id := services.Register(msg.Url)

	// Ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id": id,
	})
}

func UrlRedirect(w http.ResponseWriter, req *http.Request) {
	// Валидация
	if req.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Работа с запросом
	idStr := req.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Ресурс не найден", http.StatusNotFound)
		return
	}

	// Бизнес логика
	foundUrl := services.Find(idStr)
	if foundUrl == nil {
		http.Error(w, "Ресурс не найден", http.StatusNotFound)
		return
	}

	http.Redirect(w, req, foundUrl.GetOldUrl(), http.StatusFound)
}

// TODO: make ENV

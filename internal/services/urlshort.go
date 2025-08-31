package services

import (
	"fmt"
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const URL = "http://localhost:8080"

type UrlData struct {
	id  string
	url string
}

func NewUrlData(id, url string) UrlData {
	return UrlData{id, url}
}

// Вместо БД
var Urls = []UrlData{}

func Find(s string) *UrlData {
	for _, data := range Urls {
		if data.id == s {
			return &data
		}
	}

	return nil
}

// Получаем URL, генерируем ID и добавляем в БД
func Register(s string) string {
	id := generateId(6)

	Urls = append(Urls, NewUrlData(id, s))

	return id
}

func (u *UrlData) GetOldUrl() string {
	return u.url
}

func (u *UrlData) BuildShortUrl() string {
	url := fmt.Sprintf("%s/?id=%s", URL, u.id)

	return url
}

func generateId(length int) string {
	var sb strings.Builder
	sb.Grow(length)

	charsetLen := len(charset)

	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(charsetLen)])
	}

	return sb.String()
}

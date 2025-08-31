package services

import (
	"fmt"
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const URL = "http://localhost:8080"

func Short() string {
	url := fmt.Sprintf("%s/r/%s", URL, generateUrl(6))

	return url
}

func generateUrl(length int) string {
	var sb strings.Builder
	sb.Grow(length)

	charsetLen := len(charset)

	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(charsetLen)])
	}

	return sb.String()
}

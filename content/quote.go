package content

import (
	"encoding/json"
	"net/http"
)

type QuoteProvider struct{}

func NewQuoteProvider() *QuoteProvider {
	return &QuoteProvider{}
}

func (q *QuoteProvider) GetContent() (string, error) {
	// Используем API для цитат
	resp, err := http.Get("https://zenquotes.io/api/today")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result []struct {
		Content string `json:"q"`
		Author  string `json:"a"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result) > 0 {
		return result[0].Content + result[0].Author, nil
	}
	return "Не удалось получить цитату", nil
}

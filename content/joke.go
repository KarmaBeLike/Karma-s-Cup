package content

import (
	"encoding/json"
	"net/http"
)

type JokeProvider struct{}

func NewJokeProvider() *JokeProvider {
	return &JokeProvider{}
}

func (j *JokeProvider) GetContent() (string, error) {
	// Используем API анекдотов (можно заменить на другой источник)
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Type     string `json:"type"`
		Setup    string `json:"setup"`
		Delivery string `json:"delivery"`
		Joke     string `json:"joke"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// В зависимости от типа шутки возвращаем либо одну строку, либо setup + delivery
	if result.Type == "single" {
		return result.Joke, nil
	}
	return result.Setup + "\n" + result.Delivery, nil
}

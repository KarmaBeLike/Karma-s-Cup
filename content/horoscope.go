package content

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HoroscopeProvider struct {
	Sign string
}

func NewHoroscopeProvider(sign string) *HoroscopeProvider {
	return &HoroscopeProvider{
		Sign: sign,
	}
}

func (h *HoroscopeProvider) GetContent() (string, error) {
	url := fmt.Sprintf("https://horoscope-app-api.vercel.app/api/v1/get-horoscope/daily?sign=%s&day=TODAY", h.Sign)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			HoroscopeData string `json:"horoscope_data"`
			Date          string `json:"date"`
		} `json:"data"`
		Status  int  `json:"status"`
		Success bool `json:"success"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Data.HoroscopeData == "" {
		return "", fmt.Errorf("пустой гороскоп в ответе")
	}

	return fmt.Sprintf("%s:\n%s",
		result.Data.Date,
		result.Data.HoroscopeData), nil
}

package domain

import (
	"encoding/json"
)

type RateData struct {
	Success   string             `json:"success"`
	TimeStamp string             `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

func getRates(data []byte) map[string]float64 {
	var rates RateData
	json.Unmarshal([]byte(data), &rates)
	return rates.Rates
}

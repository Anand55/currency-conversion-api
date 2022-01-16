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

// getRates fills RateData struct instane with data from fixer api response
func getRates(data []byte) map[string]float64 {
	var rates RateData
	json.Unmarshal(data, &rates)
	return rates.Rates
}

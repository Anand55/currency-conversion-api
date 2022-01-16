package domain

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strings"
	"time"

	"github.com/Anand55/currency-conversion-api/cmd/config"
)

type ConvertResult struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Result string  `json:"result"`
}

type ConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

var (
	fixerApiUrl = `http://data.fixer.io/api/latest?access_key=%s`
)

// ConvertCurrency makes a call for fixer api and returns the currency
// exchange rate for given input request
func (c *currencyExhanger) ConvertCurrency(convertReq ConvertRequest) (ConvertResult, error) {
	// Getting fixer key from configs
	fixerKey := config.FIXER_KEY
	// Creating url with adding fixer access key
	url := fmt.Sprintf(fixerApiUrl, fixerKey)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)

	if err != nil {
		return ConvertResult{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	rdata := getRates(body)
	if convertReq.Amount == 0{
		convertReq.Amount = 1
	}
	convertedResult, err := convert(convertReq.From, convertReq.To, convertReq.Amount, rdata)
	if err != nil {
		log.Println("Error fetching converted result", err)
		return ConvertResult{}, err
	}

	return convertedResult, nil
}

// convert function exchanges currency rates of given request
func convert(from string, to string, amount float64, rdata map[string]float64) (ConvertResult, error) {
	if len(rdata) == 0 {
		return ConvertResult{}, errors.New("Empty rdata map")
	}
	rdata["EUR"] = 1
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)
	var amountConverted float64
	var res ConvertResult
	res.Amount = amount
	res.From = from
	res.To = to
	amountConverted = (rdata[to] / rdata[from]) * amount
	res.Result = fmt.Sprintf("%f in %s is equals to %f in %s", amount, from, amountConverted, to)
	return res, nil
}

package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type ConvertResult struct {
	From   string
	To     string
	Amount float64
	Result string
}

type ConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

var (
	fixerApiUrl = `http://data.fixer.io/api/latest?access_key=%s`
)

func (c *currencyExhanger) ConvertCurrency(convertReq ConvertRequest) (ConvertResult, error) {
	fixerKey := os.Getenv("FIXER_KEY")
	url := fmt.Sprintf(fixerApiUrl,fixerKey)
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

	// fmt.Println(string(body))
	rdata := getRates(body)
	fmt.Println(rdata)
	convertedResult := convert(convertReq.From, convertReq.To, convertReq.Amount, rdata)

	return convertedResult, err
}

func convert(from string, to string, amount float64, rdata map[string]float64) ConvertResult {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)
	var amountConverted float64
	var res ConvertResult
	res.Amount = amount
	res.From = from
	res.To = to
	fmt.Println("USD", rdata["USD"])
	if from == "EUR" {
		amountConverted = rdata[to] * amount
	} else {
		amountConverted = (rdata[to] / rdata[from]) * amount
	}

	res.Result = fmt.Sprintf("%f in %s is equals to %f in %s", amount, from, amountConverted, to)
	return res
}

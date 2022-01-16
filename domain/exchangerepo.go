package domain

type CurrencyExhanger interface {
	ConvertCurrency(convertReq ConvertRequest) (ConvertResult, error)
}

type currencyExhanger struct {
	convertCurr *ConvertResult
}

// NewCurrencyExhanger returns a CurrencyExhanger object
func NewCurrencyExhanger() CurrencyExhanger {
	return &currencyExhanger{
		convertCurr: &ConvertResult{},
	}
}

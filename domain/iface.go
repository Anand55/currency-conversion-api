package domain

type CurrencyExhanger interface {
	ConvertCurrency(convertReq ConvertRequest) (ConvertResult, error)
}

type currencyExhanger struct {
	convertCurr *ConvertResult
}

func NewCurrencyExhanger() CurrencyExhanger {
	return &currencyExhanger{
		convertCurr: &ConvertResult{},
	}
}

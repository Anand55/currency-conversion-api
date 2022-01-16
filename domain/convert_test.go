package domain

import (
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var server *httptest.Server

func init() {
	os.Setenv("FIXER_KEY", "d244f7d57ffda45bc7b3b39e1ae75d0d")

}
func Test_currencyExhanger_ConvertCurrency(t *testing.T) {
	// var convertCurrency *ConvertResult
	type fields struct {
		convertCurr *ConvertResult
	}
	type args struct {
		convertReq ConvertRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ConvertResult
		wantErr bool
	}{
		{
			name: "Happy test",
			fields: fields{
				convertCurr: &ConvertResult{},
			},
			args: args{
				convertReq: ConvertRequest{
					From:   "USD",
					To:     "EUR",
					Amount: 1,
				},
			},
			want: ConvertResult{
				From:   "USD",
				To:     "EUR",
				Amount: 1,
				Result: "1.000000 in USD is equals to 0.875955 in EUR",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &currencyExhanger{
				convertCurr: tt.fields.convertCurr,
			}
			got, err := c.ConvertCurrency(tt.args.convertReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("currencyExhanger.ConvertCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("currencyExhanger.ConvertCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		from   string
		to     string
		amount float64
		rdata  map[string]float64
	}
	tests := []struct {
		name string
		args args
		want ConvertResult
	}{
		{
			name: "Happy test",
			args: args{
				from:   "USD",
				to:     "EUR",
				amount: 1,
				rdata: map[string]float64{
					"USD": 1.141611,
				},
			},
			want: ConvertResult{
				From:   "USD",
				To:     "EUR",
				Amount: 1,
				Result: "1.000000 in USD is equals to 0.875955 in EUR",
			},
		},

		{
			name: "When rdata is empty",
			args: args{
				from:   "USD",
				to:     "EUR",
				amount: 1,
				rdata:  map[string]float64{},
			},
			want: ConvertResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := convert(tt.args.from, tt.args.to, tt.args.amount, tt.args.rdata); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

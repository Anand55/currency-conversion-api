package domain

import (
	"reflect"
	"testing"
)

func Test_getRates(t *testing.T) {
	var rates RateData
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want map[string]float64
	}{
		{
			name: "Happy test",
			args: args{
				data: []byte(`{"success":true,"timestamp":1642318144,"base":"EUR","date":"2022-01-16","rates":{"AED":4.193127}}`),
			},
			want: map[string]float64{
				"AED": 4.193127,
			},
		},
		{
			name: "When data is empty",
			args: args{
				data: []byte(`{}`),
			},
			want: rates.Rates,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRates(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

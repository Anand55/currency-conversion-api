package routes

import (
	"github.com/Anand55/currency-conversion-api/handler"
	"github.com/gorilla/mux"
)

func RegisterRoute(r *mux.Router) {
	r.HandleFunc("/create", handler.CreateAccount)
	r.HandleFunc("/convert", handler.ConvertCurrency)
}

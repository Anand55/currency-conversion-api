package routes

import (
	"net/http"

	"github.com/Anand55/currency-conversion-api/handler"
	"github.com/Anand55/currency-conversion-api/middleware/auth"
	"github.com/gorilla/mux"
)

// RegisterRoute is a handler function, which registers a route.
func RegisterRoute(r *mux.Router) {
	r.HandleFunc("/create", handler.CreateAccount)
	r.Handle("/convert", auth.IsAuthorized(http.HandlerFunc(handler.ConvertCurrency)))
}

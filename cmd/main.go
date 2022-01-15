package main

import (
	"log"
	"net/http"

	"github.com/Anand55/currency-conversion-api/db/redis"
	"github.com/Anand55/currency-conversion-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	redis.Init("localhost:6379", 0)
	defer redis.Close()
	routes.RegisterRoute(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

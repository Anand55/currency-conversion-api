package main

import (
	"log"
	"net/http"

	"github.com/Anand55/currency-conversion-api/cmd/config"
	"github.com/Anand55/currency-conversion-api/db/redis"
	"github.com/Anand55/currency-conversion-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server at port 8080:")
	r := mux.NewRouter()
	redis.Init(config.REDIS_ADDR, 0)
	defer redis.Close()
	routes.RegisterRoute(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Anand55/currency-conversion-api/db/redis"
	"github.com/Anand55/currency-conversion-api/domain"
	"github.com/Anand55/currency-conversion-api/middleware/auth"
)

// ConvertCurrency is a handler for /convert endpoint
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {

	c := domain.NewCurrencyExhanger()

	var convRequest domain.ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&convRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if convRequest.From == "" || convRequest.To == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convertedResult, err := c.ConvertCurrency(convRequest)
	if err != nil {
		log.Println("Error getting conversion result: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(convertedResult)
}

// CreateAccount is handler for /create endpoint.
func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var userReq auth.User
	var userResp auth.UserResponse
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userReq.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accessKey := auth.GenerateUserAccessKey(userReq.Email)

	rclient := redis.GetRedisClient()

	redisVal, _ := rclient.Get(accessKey).Int()

	if redisVal > 0 {
		w.Write([]byte("User already registered, access key: "+ accessKey))
		return
	}

	err = rclient.Set(accessKey, 1, 0).Err()
	if err != nil {
		log.Println("Error storing access key in redis: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userResp.AccessKey = accessKey
	json.NewEncoder(w).Encode(userResp)

}

package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Anand55/currency-conversion-api/db/redis"
	"github.com/Anand55/currency-conversion-api/domain"
	"github.com/Anand55/currency-conversion-api/user"
)

func ConvertCurrency(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorised User, Please provide access key"))
		return
	}
	rclient := redis.GetRedisClient()

	redisVal, _ := rclient.Get(key).Int()

	fmt.Println("Redis stored Key ", redisVal)
	if redisVal == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorised User, Please provide a valid access key"))
		return
	}

	if redisVal > 15 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorised User, API access limit exceeded"))
		return
	}
	fmt.Println("Redis updated Key ", redisVal)
	err := rclient.Set(key, redisVal+1, 0).Err()
	if err != nil {
		log.Println("Error updating access key in redis: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	c := domain.NewCurrencyExhanger()

	var convRequest domain.ConvertRequest
	err = json.NewDecoder(r.Body).Decode(&convRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if convRequest.From == "" || convRequest.To == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r1, _ := c.ConvertCurrency(convRequest)

	b, _ := json.Marshal(r1)
	fmt.Fprint(w, string(b))
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var userReq user.User
	var userResp user.UserResponse
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userReq.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accessKey := GenerateAccessKey(userReq.Email)

	rclient := redis.GetRedisClient()

	err = rclient.Set(accessKey, 1, 0).Err()
	if err != nil {
		log.Println("Error storing access key in redis: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userResp.AccessKey = accessKey
	json.NewEncoder(w).Encode(userResp)

}

func GenerateAccessKey(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

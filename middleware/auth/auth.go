package auth

import (
	"log"
	"net/http"

	"github.com/Anand55/currency-conversion-api/db/redis"
)

func IsAuthorized(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Code for the middleware...
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised User, Please provide access key"))
			return
		}
		rclient := redis.GetRedisClient()

		redisVal, _ := rclient.Get(key).Int()

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
		err := rclient.Set(key, redisVal+1, 0).Err()
		if err != nil {
			log.Println("Error updating access key in redis: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

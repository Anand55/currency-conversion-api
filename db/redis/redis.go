package redis

import (
	"log"

	"github.com/go-redis/redis"
)

type redisConfig struct {
	RedisAddr string
	RedisDbID int
}

var redisConn *redis.Client

func Init(redisAddr string, redisDbID int) {
	err := makeRedisConnection(redisAddr, redisDbID)
	if err != nil {
		log.Fatalf("Error making redis connection: %v", err)
	}
}

func makeRedisConnection(redisAddr string, redisDbID int) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",        // no password set
		DB:       redisDbID, // use default DB
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return err
	}
	redisConn = rdb
	return nil
}

func GetRedisClient() *redis.Client {
	return redisConn
}

func Close() {
	redisConn.Close()
}

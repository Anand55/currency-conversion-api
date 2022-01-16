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

// Init initializes the redis credentials and attempts to connect.
func Init(redisAddr string, redisDbID int) {
	log.Println("Initializing redis...")
	err := makeRedisConnection(redisAddr, redisDbID)
	if err != nil {
		log.Fatalf("Error making redis connection: %v", err)
	}
	log.Println("Connected to redis")
}

func makeRedisConnection(redisAddr string, redisDbID int) error {
	log.Println("Connecting to redis database")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",        // no password set
		DB:       redisDbID, // use default DB
	})

	// Checking with ping if connection is established
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("Error establishing redis connection: ", err)
		return err
	}
	redisConn = rdb
	return nil
}

// Returns the redis client object
func GetRedisClient() *redis.Client {
	return redisConn
}

// Closes the redis connection
func Close() {
	redisConn.Close()
}

package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client = nil

func GetRedisClient() *redis.Client {
	if rdb != nil {
		return rdb
	}

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rdb = redisConnection
	return rdb
}

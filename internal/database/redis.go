package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")
	return rdb
}

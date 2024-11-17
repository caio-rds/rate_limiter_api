package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

// var url = "redis://default:2Mk032uWxN5C9kTcpl48ZNgBVnreQqex@redis-10367.c15.us-east-1-2.ec2.redns.redis-cloud.com:10367"

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-10367.c15.us-east-1-2.ec2.redns.redis-cloud.com:10367",
		Password: "2Mk032uWxN5C9kTcpl48ZNgBVnreQqex",
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")
	return rdb
}

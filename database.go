package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// Redis database selected
func CreateClient(dint int) *redis.Client {

	redisDatabase := redis.NewClient(&redis.Options{
		DB:       dint,
		Addr:     "db:6379",
		Password: "",
	})

	return redisDatabase

}

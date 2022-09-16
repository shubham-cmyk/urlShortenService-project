package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// Redis database selected
func CreateClient(dint int) *redis.Client {

	redisDatabase := redis.NewClient(&redis.Options{
		DB:       dint,
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	})

	return redisDatabase

}

package database

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
		Username: "infracloud",
		Password: "",
		// Addr:     os.Getenv("DATABASE_ADDRESS"),
		// Username: os.Getenv("DATBASE_USER"),
		// Password: os.Getenv("DATABASE_PASSWORD"),
	})

	return redisDatabase

}

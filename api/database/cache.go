package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateDBClient(dbNo int) *redis.Client {
	redisOptions := redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	}
	rdb := redis.NewClient(&redisOptions)

	// Test the connection
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error()) // Log the error
	}

	return rdb
}

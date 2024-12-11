package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/shreekumar2901/url-shortener/config"
)

var Ctx = context.Background()

func CreateDBClient(dbNo int) *redis.Client {
	redisOptions := redis.Options{
		Addr:     config.Config("REDIS_ADDR"),
		Password: config.Config("REDIS_PASS"),
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

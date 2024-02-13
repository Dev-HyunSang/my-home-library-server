package db

import (
	"github.com/dev-hyunsang/my-home-library-server/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	dsn := config.GetEnv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})

	return client
}

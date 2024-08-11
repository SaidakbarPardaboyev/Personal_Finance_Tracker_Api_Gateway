package redis

import (
	"api_gateway/configs"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *configs.Config) (*redis.Client, error) {
	url := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	redisClient := *redis.NewClient(&redis.Options{
		Addr:     url,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDBNumber,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &redisClient, nil
}

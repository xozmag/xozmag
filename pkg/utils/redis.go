package utils

import (
	"delivery/configs"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg configs.Configuration) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB: 0,
	})
}
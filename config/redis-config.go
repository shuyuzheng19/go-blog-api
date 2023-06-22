package config

import (
	"github.com/go-redis/redis"
	"time"
)

var REDIS *redis.Client

func getRedis() *redis.Client {
	var redisConfig = CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:        redisConfig.Addr,
		DB:          redisConfig.DB,
		Password:    redisConfig.Password,
		PoolSize:    redisConfig.PoolSize,
		PoolTimeout: time.Duration(redisConfig.Timeout) * time.Second,
	})
	return client
}

func LoadRedisConfig() {
	REDIS = getRedis()
}

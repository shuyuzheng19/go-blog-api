package config

import (
	"github.com/go-redis/redis"
	"time"
)

var Redis *redis.Client

func getRedis() *redis.Client {

	var redisConfig = LoadRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:        redisConfig.Addr,
		DB:          redisConfig.Db,
		Password:    redisConfig.Password,
		PoolSize:    20,
		PoolTimeout: time.Second * 10,
	})

	return client
}

func InitRedis() {
	Redis = getRedis()
}

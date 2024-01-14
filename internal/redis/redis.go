package redis

import (
	"github.com/YungBenn/tech-shop-microservices/config"
	rdb "github.com/redis/go-redis/v9"
)

func NewRedis(env config.EnvVars) *rdb.Client {
	rdb := rdb.NewClient(&rdb.Options{
		Addr:     env.RedisHost,
		DB:       env.RedisDB,
	})

	return rdb
}
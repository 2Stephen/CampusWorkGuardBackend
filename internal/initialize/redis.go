package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	r := AppConfig.Redis

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("Redis Connect Error: " + err.Error())
	}
}

package utils

import (
	"CampusWorkGuardBackend/internal/initialize"
	"context"
	"time"
)

var ctx = context.Background()

// Set key with expiration
func RedisSet(key string, value interface{}, expire time.Duration) error {
	return initialize.RedisClient.Set(ctx, key, value, expire).Err()
}

// Get key
func RedisGet(key string) (string, error) {
	return initialize.RedisClient.Get(ctx, key).Result()
}

// Delete key
func RedisDel(key string) error {
	return initialize.RedisClient.Del(ctx, key).Err()
}

// Get TTL
func RedisTTL(key string) (time.Duration, error) {
	return initialize.RedisClient.TTL(ctx, key).Result()
}

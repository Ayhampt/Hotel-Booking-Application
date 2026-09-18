package cache

import (
	"context"
	"fmt"
	"time"

	redisConfig "AuthInGo/config/redis"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache() (*Cache, error) {
	client := redisConfig.ConnectToRedis()
	if client == nil {
		return nil, fmt.Errorf("failed to connect to redis")
	}

	return &Cache{client: client}, nil
}

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
	ctx := context.Background()
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
	ctx := context.Background()
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Delete(key string) error {
	ctx := context.Background()
	return c.client.Del(ctx, key).Err()
}

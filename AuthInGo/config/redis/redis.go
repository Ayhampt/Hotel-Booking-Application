package config

import (
	"context"
	"fmt"

	env "AuthInGo/config/env"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	opts, err := redis.ParseURL(env.GetString("REDIS_URL", "redis://localhost:6379"))
	if err != nil {
		fmt.Println("Failed to parse Redis URL:", err)
		return nil
	}
	client := redis.NewClient(opts)
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return nil
	}
	fmt.Println("Redis ping response:", ping)

	return client
}

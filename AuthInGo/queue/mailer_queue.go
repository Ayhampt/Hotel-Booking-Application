package queue

import (
	redisConfig "AuthInGo/config/redis"
	"fmt"

	redisqueue "github.com/asheswook/redis-queue"
)

const MAILER_QUEUE = "authMail:mailer-queue"

func NewMailerQueue() (redisqueue.Queue, error) {
	client := redisConfig.ConnectToRedis()
	if client == nil {
		return nil, fmt.Errorf("failed to connect to Redis")
	}

	cfg := redisqueue.NewConfig()
	cfg.Redis = client
	cfg.Queue.Name = MAILER_QUEUE
	cfg.Safe.AckZSetName = MAILER_QUEUE + ":ack"

	queue, err := redisqueue.NewCommonQueue(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create mailer queue: %v", err)
	}
	fmt.Println("Mailer queue created successfully")

	return queue, nil
}

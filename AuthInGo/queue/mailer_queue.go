package queue

import (
	redisConfig "AuthInGo/config/redis"
	"AuthInGo/dto"
	"fmt"

	"go.codycody31.dev/gobullmq"
)

const MAILER_QUEUE = "queue-mailer"

func NewMailerQueue() (*gobullmq.Queue[dto.MailPayload], error) {
	client := redisConfig.ConnectToRedis()
	if client == nil {
		return nil, fmt.Errorf("failed to connect to Redis")
	}

	queue, err := gobullmq.NewQueue[dto.MailPayload](MAILER_QUEUE, client, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create mailer queue: %v", err)
	}
	fmt.Println("Mailer queue created successfully")

	return queue, nil
}

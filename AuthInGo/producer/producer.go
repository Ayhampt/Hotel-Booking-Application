package producer

import (
	"context"

	dto "AuthInGo/dto"
	queue "AuthInGo/queue"
	"fmt"
)

func PushToQueue(payload dto.MailPayload) error {
	mailerQueue, err := queue.NewMailerQueue()
	if err != nil {
		return fmt.Errorf("failed to initialize mailer queue: %w", err)
	}
	_, err = mailerQueue.Add(context.Background(), "payload-mail", payload)
	if err != nil {
		return fmt.Errorf("failed to push mail payload to queue: %w", err)
	}
	fmt.Println("Pushed the payload to redis")

	return nil
}

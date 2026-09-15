package producer

import (
	"encoding/json"

	queue "AuthInGo/queue"

	"fmt"
	dto "AuthInGo/dto"
)

func PushToQueue(payload dto.MailPayload) error {
	mailerQueue, err := queue.NewMailerQueue()
	if err != nil {
		return fmt.Errorf("failed to initialize mailer queue: %w", err)
	}
	if err != nil {
		return fmt.Errorf("failed to marshal mail payload: %w", err)
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal mail payload: %w", err)
	}

	err = mailerQueue.Push(string(payloadJSON))
	if err != nil {
		return fmt.Errorf("failed to push mail payload to queue: %w", err)
	}
	fmt.Println("Pushed the payload to redis")

	return nil
}

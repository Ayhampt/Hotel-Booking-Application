package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	redisConfig "AuthInGo/config/redis"
	dto "AuthInGo/dto"
	pro "AuthInGo/producer"
	"fmt"
)

func main() {
	config.Load()
	port := config.GetString("PORT", ":8080")

	cfg := app.NewConfig(port)
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()
	redisConfig.ConnectToRedis()
	if err := pro.PushToQueue(dto.MailPayload{
		To:      "user@example",
		Subject: "Test Email",
		Body:    "This is a test email body.",
		Token:   "sample-token",
	}); err != nil {
		fmt.Println("Failed to push mail payload:", err)
	}

	app.Run()
}

package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	redisConfig "AuthInGo/config/redis"
)

func main() {
	config.Load()
	port := config.GetString("PORT", ":8080")

	cfg := app.NewConfig(port)
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()
	redisConfig.ConnectToRedis()

	app.Run()
}

package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
)

func main() {
	config.Load()
	port := config.GetString("PORT", ":8080")

	cfg := app.NewConfig(port)
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()


	app.Run()
}
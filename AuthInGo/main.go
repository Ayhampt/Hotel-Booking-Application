package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)

func main() {
	config.Load()
	port := config.GetString("PORT", ":8080")

	cfg := app.NewConfig(port)
	app := app.NewApplication(cfg)

	app.Run()
}
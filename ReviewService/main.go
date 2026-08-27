package main

import (
	"ReviewService/app"
	config "ReviewService/config/env"
)

func main() {
	config.Load()
	port := config.GetString("PORT", ":3008")
	cfg := app.NewConfig(port)
	app := app.NewApplication(cfg)

	app.Run()

}

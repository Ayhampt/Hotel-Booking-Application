package app

import (
	dbConfig "ReviewService/config/db"
	"ReviewService/routers"
	"fmt"
	"net/http"
	"time"
)

type config struct {
	Addr string
}

type Application struct {
	config config
}

func NewConfig(addr string) config {
	return config{
		Addr: addr,
	}
}

func NewApplication(cfg config) *Application {
	return &Application{
		config: cfg,
	}
}

func (app *Application) Run() error {
	_, err := dbConfig.SetupDB()
	if err != nil {
		return fmt.Errorf("failed to setup database: %v", err)
	}

	server := &http.Server{
		Addr:         app.config.Addr,
		Handler:      routers.SetupRouter(nil),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Server is running on %s\n", app.config.Addr)
	return server.ListenAndServe()
}

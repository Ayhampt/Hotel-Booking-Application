package app

import (
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/routers"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
	Store db.Storage
}


func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store: *db.NewStorage(),
	}
}


func (app *Application) Run() error {
	ur:= db.NewUserRepository()
	us:= services.NewUserService(ur)
	uc:= controllers.NewUserController(us)
	uRouter:=routers.NewUserRouter(uc)

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: routers.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is starting in port:",app.Config.Addr)
	return server.ListenAndServe()

}


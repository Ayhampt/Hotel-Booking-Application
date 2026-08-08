package app

import (
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/routers"
	"AuthInGo/services"
	dbConfig "AuthInGo/config/db"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}


func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}


func (app *Application) Run() error {

	db,err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Error in setting up database",err)
		return err
	}
	ur:= repo.NewUserRepository(db)
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


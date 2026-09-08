package app

import (
	dbConfig "AuthInGo/config/db"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
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

	db, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Error in setting up database", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	rpr := repo.NewRolePermissionRepository(db)
	pr := repo.NewPermissionRepository(db)
	urr := repo.NewUserRoleRepository(db)
	us := services.NewUserService(ur)
	rs := services.NewRoleService(rr, rpr, urr, pr)
	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	uRouter := routers.NewUserRouter(uc)
	rRouter := routers.NewRoleRouter(rc) 

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      routers.SetupRouter(uRouter, rRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is starting in port:", app.Config.Addr)
	return server.ListenAndServe()

}

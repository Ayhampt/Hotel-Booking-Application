package routers

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger) //build-in chi middleware for logging
	chiRouter.Use(middlewares.RateLimiterMiddleware)

	chiRouter.Get("/ping",controllers.PingHandler)
	chiRouter.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.com", "/fakestoreservice"))

	UserRouter.Register(chiRouter)



	return chiRouter

}
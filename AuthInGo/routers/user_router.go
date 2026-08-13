package routers

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)


type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}

}

func (ur *UserRouter) Register(r chi.Router) {
	r.With(middlewares.CreateUserRequestValidator).Post("/signup",ur.userController.CreateUser)
	r.Get("/profile",ur.userController.GetUserById)
	r.With(middlewares.LoginRequestValidator).Post("/login",ur.userController.LoginUser)
}

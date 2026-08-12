package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter,r *http.Request) {
	uc.UserService.CreateUser()
	w.Write([]byte("User registration endpoint"))
}

func (uc *UserController) GetUserById(w http.ResponseWriter,r *http.Request) {
	uc.UserService.GetUserById()
	w.Write([]byte("Get user by id endpoint"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Login user endpoint")

	var payload dto.LoginUserRequestDto

	if jsonErr := utils.ReadJsonBody(r,&payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Something went wrong while logging in",jsonErr)
		return
	}

	if validationErr := utils.Validator.Struct(payload);validationErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid Input Data",validationErr)
		return
	}

	jwtToken,err := uc.UserService.LoginUser(&payload)
	fmt.Println("JWT Token:", jwtToken)
	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed To Login",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User Logged In Successfully",jwtToken)

}
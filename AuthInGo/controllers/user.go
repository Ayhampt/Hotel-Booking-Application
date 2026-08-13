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

func (uc *UserController) CreateUser(w http.ResponseWriter,r *http.Request) {
	payload := r.Context().Value("CreateUserPayload").(dto.CreateUserRequestDto)
	user,err := uc.UserService.CreateUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed To Create User",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User Created Successfully",user)
}

func (uc *UserController) GetUserById(w http.ResponseWriter,r *http.Request) {
	userId := r.URL.Query().Get("id")
	if userId == "" {
		userId = r.Context().Value("userID").(string)
	}

	fmt.Println("User ID from context or query:", userId)

	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}
	user, err := uc.UserService.GetUserById(userId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %d not found", userId))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully:", user)

}

func (uc *UserController) LoginUser(w http.ResponseWriter,r *http.Request) {
	payload := r.Context().Value("LoginUserPayload").(dto.LoginUserRequestDto)

	jwtToken,err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Failed To Login",err)
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User Logged In Successfully",jwtToken)

}
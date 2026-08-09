package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	CreateUser() error
	GetUserById() error
	LoginUser() (string,error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "test2026"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password", err)
		return err
	}
	u.userRepository.Create("ayhaj", "ayhsa@gmail.com", hashedPassword)
	return nil
}

func (u *UserServiceImpl) LoginUser() (string,error) {
	email := "ayha@gmail.com"
	password := "test2026"

	user,err := u.userRepository.GetByEmail(email)

	if err != nil{
		fmt.Println("Error in fetching user by email:",err)
		return "",err
	}
	if user == nil {
		fmt.Println("User not found with given email")
		return "",fmt.Errorf("User not found with given email %s",email)
	}

	isPasswordValid := utils.CheckPasswordHash(password,user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", fmt.Errorf("invalid credentials")
	}

	payload := jwt.MapClaims {
		"email":user.Email,
		"id":user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	tokenString,err := token.SignedString([]byte(env.GetString("JWT_SECRET","TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:",err)
		return "",err
	}

	fmt.Println("Token generated successfully:",tokenString)
	return tokenString,nil
}


func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Getting user by Id")
	u.userRepository.GetById()
	return nil
}

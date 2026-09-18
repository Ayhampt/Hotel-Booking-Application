package services

import (
	"AuthInGo/cache"
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	dto "AuthInGo/dto"
	"AuthInGo/models"
	pro "AuthInGo/producer"
	"AuthInGo/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	CreateUser(payload *dto.CreateUserRequestDto) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDto) (string, error)
	VerifyUser(token string) error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDto) (*models.User, error) {
	fmt.Println("Creating user in UserService")
	password := payload.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password", err)
		return nil, err
	}
	user, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error creating user in repository", err)
		return nil, err
	}
	token, err := utils.GenerateVerifyToken(32)
	if err != nil {
		fmt.Println("Error generating verification token:", err)
		return nil, err
	}
	cache, err := cache.NewCache()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil, err
	}
	err = cache.Set(token, fmt.Sprintf("%d", user.Id), 24*time.Hour)
	if err != nil {
		fmt.Println("Error setting token in Redis:", err)
		return nil, err
	}
	if err := pro.PushToQueue(dto.MailPayload{
		To:         user.Email,
		Subject:    "Verify your email",
		TemplateID: "verify-email",
		Params: dto.MailParams{
			Token:           token,
			VerificationURL: fmt.Sprintf("%s/verify?token=%s", env.GetString("FRONTEND_URL", "http://localhost:3001"), token),
		},
	}); err != nil {
		fmt.Println("Failed to push mail payload:", err)
	}
	fmt.Println("User created successfully with ID:", user.Id)
	return user, nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDto) (string, error) {

	user, err := u.userRepository.GetByEmail(payload.Email)

	if err != nil {
		fmt.Println("Error in fetching user by email:", err)
		return "", err
	}
	if user == nil {
		fmt.Println("User not found with given email")
		return "", fmt.Errorf("User not found with given email %s", payload.Email)
	}

	isPasswordValid := utils.CheckPasswordHash(payload.Password, user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", fmt.Errorf("invalid credentials")
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("Token generated successfully:", tokenString)
	return tokenString, nil
}

func (u *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	fmt.Println("Getting user by Id")
	user, err := u.userRepository.GetById(id)
	if err != nil {
		fmt.Println("Error fetching user by Id:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) VerifyUser(token string) error {
	cache, err := cache.NewCache()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return err
	}
	userId, err := cache.Get(token)
	if err != nil {
		fmt.Println("Error getting token from Redis:", err)
		return err
	}
	if userId == "" {
		fmt.Println("Token not found or expired")
		return fmt.Errorf("token not found or expired")
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("Error parsing user ID:", err)
		return err
	}

	err = u.userRepository.UpdateIsVerified(userIdInt)
	if err != nil {
		fmt.Println("Error verifying user:", err)
		return err
	}

	err = cache.Delete(token)
	if err != nil {
		fmt.Println("Error deleting token from Redis:", err)
		return err
	}

	fmt.Println("User verified successfully with ID:", userId)
	return nil
}

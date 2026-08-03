package db

import (
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct {
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{

	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("creating user in userRepository")
	return nil
}
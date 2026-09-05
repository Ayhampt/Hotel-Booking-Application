package db

import (
	"database/sql"
	"AuthInGo/models"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func (r *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	return nil, nil
}


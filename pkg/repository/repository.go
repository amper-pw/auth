package repository

import (
	"github.com/amper-pw/auth/pkg/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(username, password string) (*model.User, error)
	FindUserByUsernameAndPassword(username, password string) (*model.User, error)
}

type Repository struct {
	UserRepository UserRepositoryInterface
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}

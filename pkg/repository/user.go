package repository

import (
	"errors"
	"github.com/amper-pw/auth/pkg/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) FindUserByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("username= ? AND password= ?", username, password).First(&user).Error; err != nil {
		err := errors.New("username or password is invalid")
		return &user, err
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(username, password string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,
	}
	if err := r.db.Create(user).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

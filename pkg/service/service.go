package service

import (
	"github.com/amper-pw/auth/pkg/model"
	"github.com/amper-pw/auth/pkg/repository"
)

type AuthServiceInterface interface {
	RegisterUser(username, password string) (*model.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Service struct {
	AuthService AuthServiceInterface
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		AuthService: BuildAuthService(repos.UserRepository),
	}
}

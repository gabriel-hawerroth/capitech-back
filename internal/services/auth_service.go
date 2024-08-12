package services

import (
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
)

type AuthService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(UserRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: UserRepository,
	}
}

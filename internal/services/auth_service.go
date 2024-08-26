package services

import (
	"errors"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
	"golang.org/x/crypto/bcrypt"
)

const errorUserAlreadyExists = "this email is already in use"

type AuthService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(UserRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: UserRepository,
	}
}

func (s *AuthService) CreateNewUser(dto dto.CreateUserDTO) error {
	exists, err := s.UserRepository.ExistsByEmail(dto.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(errorUserAlreadyExists)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dto.Password = string(hash)

	err = s.UserRepository.CreateNewUser(dto)
	return err
}

package services

import (
	"errors"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
	"golang.org/x/crypto/bcrypt"
)

var (
	errUserAlreadyExists = errors.New("this email is already in use")
	errInvalidPasswod    = errors.New("invalid password")
	ErrGeneratingToken   = errors.New("error generating token")
)

type AuthService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(UserRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: UserRepository,
	}
}

func (s *AuthService) DoLogin(email, password string) (string, error) {
	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errInvalidPasswod
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return "", ErrGeneratingToken
	}

	return token, nil
}

func (s *AuthService) CreateNewUser(dto dto.CreateUserDTO) error {
	exists, err := s.UserRepository.ExistsByEmail(dto.Email)
	if err != nil {
		return err
	}
	if exists {
		return errUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dto.Password = string(hash)

	err = s.UserRepository.CreateNewUser(dto)
	return err
}

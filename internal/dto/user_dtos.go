package dto

import (
	"errors"
	"net/mail"
	"unicode"
)

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *CreateUserDTO) Validate() error {
	if dto.Name == "" {
		return errors.New("name is required")
	}
	if dto.Email == "" {
		return errors.New("email is required")
	}
	if dto.Password == "" {
		return errors.New("password is required")
	}

	if !isValidEmail(dto.Email) {
		return errors.New("invalid email")
	}

	if !isValidPassword(dto.Password) {
		return errors.New("invalid password")
	}

	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidPassword(password string) bool {
	letters := 0
	upper := false
	lower := false
	number := false
	special := false

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}

		letters++
	}

	return letters >= 8 && upper && lower && number && special
}

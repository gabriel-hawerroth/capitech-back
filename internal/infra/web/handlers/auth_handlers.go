package handlers

import (
	"net/http"

	"github.com/gabriel-hawerroth/capitech-back/internal/services"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(AuthService services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: AuthService,
	}
}

func (h *CategoryHandler) DoLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *CategoryHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

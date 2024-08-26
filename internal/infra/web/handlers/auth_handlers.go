package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
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

func (h *AuthHandler) DoLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *AuthHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, errorDecodingRequestBody, http.StatusInternalServerError)
		return
	}

	err := h.AuthService.CreateNewUser(dto)
	if err != nil {
		setHttpError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

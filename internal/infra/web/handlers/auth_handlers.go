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
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	res, err := h.AuthService.DoLogin(email, password)
	if err != nil {
		if err == services.ErrGeneratingToken {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "bad credentials", http.StatusUnauthorized)
		return
	}

	setJsonContentType(w)
	if err := json.NewEncoder(w).Encode(map[string]string{"token": res}); err != nil {
		http.Error(w, errorEncodingResponse, http.StatusInternalServerError)
	}
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

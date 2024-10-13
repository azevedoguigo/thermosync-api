package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/thermosync-api/internal/contract"
	"github.com/azevedoguigo/thermosync-api/internal/service"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(service service.UserService) *AuthHandler {
	return &AuthHandler{userService: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials contract.LoginDTO
	json.NewDecoder(r.Body).Decode(&credentials)

	token, err := h.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

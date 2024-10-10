package handler

import (
	"encoding/json"
	"net/http"

	"github.com/azevedoguigo/thermosync-api/internal/contract"
	"github.com/azevedoguigo/thermosync-api/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto contract.NewUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload!", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

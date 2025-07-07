package handler

import (
	"encoding/json"
	"net/http"

	"github.com/siriramhazam/budget-authen/internal/service"
	"github.com/siriramhazam/budget-authen/internal/utils"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		svc: service.NewAuthService(),
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	if body.Username == "" {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials", "Missing field: Username")
		return
	}

	if body.Password == "" {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials", "Missing field: Password")
		return
	}

	token, err := h.svc.GenerateToken(body.Username)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Could not genarate token", err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/siriramhazam/budget-authen/internal/service"
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
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if body.Username == "" {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if body.Password == "" {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.svc.GenerateToken(body.Username)
	if err != nil {
		fmt.Println("Error generating token:", err)
		http.Error(w, "could not genarate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}

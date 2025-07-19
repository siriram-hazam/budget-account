package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey")

var (
	ErrInvalidUsername = errors.New("username is required")
	ErrInvalidPassword = errors.New("password is required")
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	if username == "" {
		return "", ErrInvalidUsername
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

package service

import (
	"strings"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	authService := NewAuthService()

	username := "testuser"
	token, err := authService.GenerateToken(username)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatalf("Expected a token, got an empty string")
	}

	if !strings.Contains(token, ".") {
		t.Fatalf("Expected token to be a valid JWT, got %v", token)
	}
}

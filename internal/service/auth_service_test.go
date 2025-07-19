package service

import (
	"context"
	"strings"
	"testing"

	pb "github.com/siriramhazam/budget-authen/grpc-auth/proto"
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
		t.Errorf("Expected token to be a valid JWT, got %v", token)
	}
}

func TestGenerateToken_EmptyUsername(t *testing.T) {
	authService := NewAuthService()

	token, err := authService.GenerateToken("")

	if err == nil {
		t.Errorf("Expected error for empty username, got nil")
	}
	if token != "" {
		t.Errorf("Expected empty token for empty username, got %v", token)
	}
}

type mockAuthServer struct {
	pb.UnimplementedAuthServiceServer
	authService *AuthService
}

func newMockAuthServer(authService *AuthService) *mockAuthServer {
	return &mockAuthServer{authService: authService}
}

func (s *mockAuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Username == "" {
		return nil, ErrInvalidUsername
	}
	if req.Password == "" {
		return nil, ErrInvalidPassword
	}
	token, err := s.authService.GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func TestAuthServer_Login(t *testing.T) {
	authService := NewAuthService()
	server := newMockAuthServer(authService)

	tests := []struct {
		name     string
		username string
		password string
		wantErr  bool
	}{

		{"valid", "testuser", "testpass", false},
		{"missing username", "", "testpass", true},
		{"missing password", "testuser", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pb.LoginRequest{
				Username: tt.username,
				Password: tt.password,
			}
			resp, err := server.Login(context.Background(), req)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, gor nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if resp == nil || resp.Token == "" {
					t.Errorf("expected token, got nil or empty")
				}
			}
		})
	}
}

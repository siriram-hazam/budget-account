package main

import (
	"context"
	"log"
	"net"

	pb "github.com/siriramhazam/budget-authen/grpc-auth/proto"
	"github.com/siriramhazam/budget-authen/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	authService *service.AuthService
}

func NewAuthServer(authService *service.AuthService) *AuthServer {
	return &AuthServer{
		authService: authService,
	}
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.authService.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	authService := service.NewAuthService()

	authServer := NewAuthServer(authService)
	pb.RegisterAuthServiceServer(s, authServer)

	log.Println("gRPC server starting at :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// Handler Http
// func main() {
// 	mux := http.NewServeMux()
// 	authHandler := handler.NewAuthHandler()
// 	mux.HandleFunc("/auth", authHandler.Login)

// 	log.Println("Server stating at :8080")
// 	if err := http.ListenAndServe(":8080", mux); err != nil {
// 		log.Fatal(err)
// 	}
// }

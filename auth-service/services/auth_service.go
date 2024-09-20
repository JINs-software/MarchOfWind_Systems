package services

import (
	"JINs-software/MOW_SYSTEMS/auth-service/grpc"
	"JINs-software/MOW_SYSTEMS/auth-service/models"
	"context"
	"errors"
)

type AuthService struct {
	grpc.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	user, exists := models.Users[req.Username]
	if !exists || user.Password != req.Password {
		return nil, errors.New("invalid credentials")
	}

	return &grpc.LoginResponse{
		Token: user.Token,
	}, nil
}

package handler

import (
	"context"

	"github.com/CatGitBon/auth-service/pkg/auth"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) AuthenticateUser(ctx context.Context, req *auth.RequestType) (*auth.ResponseType, error) {
	// Здесь будет ваша логика аутентификации
	// Пока что возвращаем заглушку
	return &auth.ResponseType{
		Response: true,
	}, nil
}

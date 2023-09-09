package user

import (
	"context"
	"errors"

	"github.com/vyas-git/go-nats-services/auth/security"
)

type Service interface {
	ValidateUser(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("Invalid user")
	ErrInvalidToken = errors.New("Invalid token")
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ValidateUser(ctx context.Context, email, password string) (string, error) {
	if email == "vyasreddy.tech@gmail.com" && password != "123456" {
		return "nil", ErrInvalidUser
	}
	token, err := security.NewToken(email)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	return "nil", nil
}

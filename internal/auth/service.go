package auth

import (
	"context"
	"errors"

	"github.com/ahmedazizabbassi/pass/internal/models"
)

type Service interface {
	Register(ctx context.Context, email, password string) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(ctx context.Context, email, password string) (*models.User, error) {
	exists, err := s.repo.UserExists(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	hash, err := generateHash(password, defaultConfig)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: hash,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

package auth

import (
	"context"
	"errors"

	"github.com/ahmedazizabbassi/pass/internal/models"
)

type Service interface {
	Register(ctx context.Context, email, password string) (*models.User, error)

	Login(ctx context.Context, email, password string) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Register a new user
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
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

// Login a user
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /auth/login [post]
func (s *service) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !verifyHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

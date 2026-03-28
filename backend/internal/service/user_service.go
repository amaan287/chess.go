package service

import (
	"context"

	"github.com/amaan287/chess-backend/internal/models"
	"github.com/amaan287/chess-backend/internal/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	return s.repo.CreateUser(ctx, user)
}

package ports

import (
	"context"

	"github.com/amaan287/chess-backend/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUserByIdentifier(ctx context.Context, identifier string) (models.User, error)
	ListUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

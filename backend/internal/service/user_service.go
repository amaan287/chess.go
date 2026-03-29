package service

import (
	"context"
	"errors"
	"strings"

	"github.com/amaan287/chess-backend/internal/auth"
	"github.com/amaan287/chess-backend/internal/models"
	"github.com/amaan287/chess-backend/internal/ports"
	"github.com/amaan287/chess-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidRefreshToken = errors.New("invalid refresh token")

type UserService struct {
	repo         ports.UserRepository
	tokenManager *auth.Manager
}

func NewUserService(repo ports.UserRepository, tokenManager *auth.Manager) *UserService {
	return &UserService{
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	if user == nil {
		return models.User{}, errors.New("user is nil")
	}

	if strings.TrimSpace(user.Password) == "" {
		return models.User{}, errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s *UserService) Login(ctx context.Context, identifier, password string) (models.User, string, string, error) {
	identifier = strings.TrimSpace(identifier)
	if identifier == "" || strings.TrimSpace(password) == "" {
		return models.User{}, "", "", ErrInvalidCredentials
	}

	user, err := s.repo.GetUserByIdentifier(ctx, identifier)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return models.User{}, "", "", ErrInvalidCredentials
		}
		return models.User{}, "", "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, "", "", ErrInvalidCredentials
	}

	if s.tokenManager == nil {
		return models.User{}, "", "", errors.New("token manager is nil")
	}

	accessToken, refreshToken, err := s.tokenManager.GenerateTokenPair(user.ID)
	if err != nil {
		return models.User{}, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (s *UserService) RefreshAccessToken(refreshToken string) (string, error) {
	if s.tokenManager == nil {
		return "", errors.New("token manager is nil")
	}

	claims, err := s.tokenManager.ParseToken(refreshToken, auth.TokenTypeRefresh)
	if err != nil {
		return "", ErrInvalidRefreshToken
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(claims.UserID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

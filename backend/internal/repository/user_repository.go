package repository

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/amaan287/chess-backend/constants"
	"github.com/amaan287/chess-backend/internal/models"
	"github.com/amaan287/chess-backend/internal/ports"
)

type PostgresUserRepository struct {
	db *sql.DB
}

var _ ports.UserRepository = (*PostgresUserRepository)(nil)

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	if user == nil {
		return models.User{}, constants.ErrNilUser
	}

	if user.ID == "" {
		id, err := generateID()
		if err != nil {
			return models.User{}, err
		}
		user.ID = id
	}

	if user.Rating == 0 {
		user.Rating = 1200
	}

	now := time.Now().UTC()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.LastLogin.IsZero() {
		user.LastLogin = now
	}

	const query = `
		INSERT INTO users (id, username, name, email, password, rating, created_at, last_login)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, username, name, email, password, rating, created_at, last_login
	`

	createdUser, err := scanUser(r.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Name,
		user.Email,
		user.Password,
		user.Rating,
		user.CreatedAt,
		user.LastLogin,
	))
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id string) (models.User, error) {
	const query = `
		SELECT id, username, name, email, password, rating, created_at, last_login
		FROM users
		WHERE id = $1
	`

	user, err := scanUser(r.db.QueryRowContext(ctx, query, id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, constants.ErrUserNotFound
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *PostgresUserRepository) GetUserByIdentifier(ctx context.Context, identifier string) (models.User, error) {
	const query = `
		SELECT id, username, name, email, password, rating, created_at, last_login
		FROM users
		WHERE email = $1 OR username = $1
		LIMIT 1
	`

	user, err := scanUser(r.db.QueryRowContext(ctx, query, identifier))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, constants.ErrUserNotFound
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *PostgresUserRepository) ListUsers(ctx context.Context) ([]models.User, error) {
	const query = `
		SELECT id, username, name, email, password, rating, created_at, last_login
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Rating,
			&user.CreatedAt,
			&user.LastLogin,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *models.User) (models.User, error) {
	if user == nil {
		return models.User{}, constants.ErrNilUser
	}

	if user.ID == "" {
		return models.User{}, constants.ErrUserIdRequired
	}

	if user.LastLogin.IsZero() {
		user.LastLogin = time.Now().UTC()
	}

	const query = `
		UPDATE users
		SET username = $2, name = $3, email = $4, password = $5, rating = $6, last_login = $7
		WHERE id = $1
		RETURNING id, username, name, email, password, rating, created_at, last_login
	`

	updatedUser, err := scanUser(r.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Name,
		user.Email,
		user.Password,
		user.Rating,
		user.LastLogin,
	))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, constants.ErrUserNotFound
		}
		return models.User{}, err
	}

	return updatedUser, nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id string) error {
	const query = `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return constants.ErrUserNotFound
	}

	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanUser(row scanner) (models.User, error) {
	user := models.User{}
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Rating,
		&user.CreatedAt,
		&user.LastLogin,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func generateID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

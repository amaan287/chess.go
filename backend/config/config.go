package config

import (
	"database/sql"
	"fmt"

	"github.com/amaan287/chess-backend/constants"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	envConfig, err := constants.GetEnv()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		envConfig.Host,
		envConfig.DBPort,
		envConfig.DBUser,
		envConfig.DBPassword,
		envConfig.DBName,
		envConfig.DBSSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) DB() *sql.DB {
	return s.db
}

func (s *PostgresStore) Close() error {
	if s == nil || s.db == nil {
		return nil
	}

	return s.db.Close()
}

package migrations

import (
	"context"
	"database/sql"
)

func Run(ctx context.Context, db *sql.DB) error {
	return createUsersTable(ctx, db)
}

func createUsersTable(ctx context.Context, db *sql.DB) error {
	const query = `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			rating INTEGER NOT NULL DEFAULT 1200 CHECK (rating >= 0),
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			last_login TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`

	_, err := db.ExecContext(ctx, query)
	return err
}

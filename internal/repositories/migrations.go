// Package repositories contains the methods that interact with the database
package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RunMigrations creates the table if it does not exist
func RunMigrations(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS courses (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description VARCHAR(255) NOT NULL
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to run migrations to create or load the table: %w", err)
	}

	return nil
}

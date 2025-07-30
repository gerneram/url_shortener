package migrations

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	const op = "storage.migrations.RunMigrations"

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS url(
    				id serial PRIMARY KEY,
    				alias TEXT NOT NULL UNIQUE,
    				url TEXT NOT NULL);
    				`)
	if err != nil {
		return fmt.Errorf("%s: create table url: %w", op, err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);`)
	if err != nil {
		return fmt.Errorf("%s: create index: %w", op, err)
	}
	return nil
}

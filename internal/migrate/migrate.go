// Package migrate provides embedded to service SQL migrations
package migrate

import (
	"embed"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	//go:embed migrations/*
	dbMigrations embed.FS
	migrations   migrate.EmbedFileSystemMigrationSource
)

func init() {
	migrations = migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}
}

// Up - migrate the database up
func Up(db *sqlx.DB) error {
	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	return err
}

// Down - revert latest migrations
func Down(db *sqlx.DB) error {
	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Down)
	return err
}

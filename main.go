// Package main entry of the service
package main

import (
	"context"
	"log"

	"github.com/Striboh/viter/internal/cli"
	"github.com/Striboh/viter/internal/config"
	"github.com/Striboh/viter/internal/migrate"
	"github.com/Striboh/viter/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	config := config.FromEnv()
	db, err := sqlx.Open("postgres", config.PostgresURL)
	if err != nil {
		log.Fatalf("Failed initializing connection to Postgres DB: %s", err)
	}

	switch cli.GetMigrateMode() {
	case cli.MigrateModeUp:
		err = migrate.Up(db)
	case cli.MigrateModeDown:
		err = migrate.Down(db)
	}

	if err != nil {
		log.Fatalf("Failed migrating DB: %s", err)
	}

	ctx := context.Background()
	if err := service.Run(ctx, config, db); err != nil {
		panic(err)
	}
}

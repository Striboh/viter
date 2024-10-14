// Package service provides API and (in future) background processes.
package service

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Striboh/viter/internal/config"
	"github.com/Striboh/viter/internal/service/api"
	"github.com/jmoiron/sqlx"
)

// Run - start the API and all background services
func Run(ctx context.Context, config config.Config, db *sqlx.DB) error {
	slog.Info("Starting server", "addr", config.ListenAddr)
	// create a type that satisfies the `api.ServerInterface`,
	// which contains an implementation of every operation from
	// the generated code
	server := api.NewServer(config, db)

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    config.ListenAddr,
		// TODO(Velnbur): make this configurable later
		ReadHeaderTimeout: time.Second * 15,
	}

	// And we serve HTTP until the world ends.
	return s.ListenAndServe()
}

// Package api provides implementation of generated from openapi.yaml
// server.
package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Striboh/viter/internal/config"
	"github.com/jmoiron/sqlx"
)

// Server - implements the API from openapi.yaml
type Server struct {
	config config.Config
	db     *sqlx.DB
}

// Compile time check hat Server implements ServerInterface
var _ ServerInterface = (*Server)(nil)

// NewServer - return Server
func NewServer(config config.Config, db *sqlx.DB) Server {
	return Server{config: config, db: db}
}

// CreateProfile - (POST /profiles) handler
func (Server) CreateProfile(w http.ResponseWriter, r *http.Request) {
	resp := Profile{
		Id: 0,
	}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error("failed writing json to response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// ShowProfileByID - (GET /profiles/{profileId}) handler
func (Server) ShowProfileByID(w http.ResponseWriter, r *http.Request, profileID string) {
	resp := Profile{
		Id:   0,
		Name: "",
	}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error("failed writing json to response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

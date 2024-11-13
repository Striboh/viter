// Package api provides implementation of generated from openapi.yaml
// server.
package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Striboh/viter/internal/config"
	"github.com/Striboh/viter/internal/models"
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
func (s Server) CreateProfile(w http.ResponseWriter, r *http.Request) {
	data_jsonstruct := Profile{}
	err := json.NewDecoder(r.Body).Decode(&data_jsonstruct)
	if err != nil {
		slog.Error("failed reading json request", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data_sqlstruct := models.Profile{
		FirstName: data_jsonstruct.FirstName,
		LastName:  data_jsonstruct.LastName,
		Email:     data_jsonstruct.Email,
		Phone:     data_jsonstruct.Phone,
		Roles:     data_jsonstruct.Roles,
	}
	id_str, err := models.CreateProfile(s.db, data_sqlstruct)
	if err != nil {
		slog.Error("failed to create profile", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success := true
	resp := ProfileResponse{
		Success: &success,
		Error:   nil,
		Id:      &id_str,
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error("failed writing json to response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// ShowProfileByID - (GET /profiles/{profileId}) handler
func (s Server) ShowProfileByID(w http.ResponseWriter, r *http.Request, profileID string) {
	resp_sqlstruct, err := models.GetProfile(s.db, profileID)
	if err != nil {
		slog.Error("failed getting profile by id", slog.Any("err", err))
	}
	resp_jsonstruct := Profile{
		Email:     resp_sqlstruct.Email,
		Phone:     resp_sqlstruct.Phone,
		Roles:     resp_sqlstruct.Roles,
		FirstName: resp_sqlstruct.FirstName,
		LastName:  resp_sqlstruct.LastName,
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp_jsonstruct)
	if err != nil {
		slog.Error("failed writing json to response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// UpdateProfileByID implements ServerInterface.
func (s Server) UpdateProfileByID(w http.ResponseWriter, r *http.Request, profileID string) {
	data_jsonstruct := Profile{}
	err := json.NewDecoder(r.Body).Decode(&data_jsonstruct)
	if err != nil {
		slog.Error("failed reading json request", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data_sqlstruct := models.Profile{
		FirstName: data_jsonstruct.FirstName,
		LastName:  data_jsonstruct.LastName,
		Email:     data_jsonstruct.Email,
		Phone:     data_jsonstruct.Phone,
		Roles:     data_jsonstruct.Roles,
	}

	err = models.UpdateProfile(s.db, data_sqlstruct, profileID)
	if err != nil {
		slog.Error("failed updating profile", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success := true
	resp := ProfileUpdateResponse{
		Success: &success,
		Error:   nil,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error("failed writing json to response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// CreateCategory implements ServerInterface.
func (s Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateOrUpdateProduct implements ServerInterface.
func (s Server) CreateOrUpdateProduct(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetApiToken implements ServerInterface.
func (s Server) GetApiToken(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetCategoryByID implements ServerInterface.
func (s Server) GetCategoryByID(w http.ResponseWriter, r *http.Request, categoryID string) {
	panic("unimplemented")
}

// GetCategoryTree implements ServerInterface.
func (s Server) GetCategoryTree(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

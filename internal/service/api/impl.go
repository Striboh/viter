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
	dataJSONstruct := Profile{}
	err := json.NewDecoder(r.Body).Decode(&dataJSONstruct)
	if err != nil {
		slog.Error("failed reading json request", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataSQLstruct := models.Profile{
		FirstName: dataJSONstruct.FirstName,
		LastName:  dataJSONstruct.LastName,
		Email:     dataJSONstruct.Email,
		Phone:     dataJSONstruct.Phone,
		Roles:     dataJSONstruct.Roles,
	}
	idStr, err := models.CreateProfile(s.db, dataSQLstruct)
	if err != nil {
		slog.Error("failed to create profile", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success := true
	resp := ProfileResponse{
		Success: &success,
		Error:   nil,
		Id:      &idStr,
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
	respSQLstruct, err := models.GetProfile(s.db, profileID)
	if err != nil {
		slog.Error("failed getting profile by id", slog.Any("err", err))
	}
	respJSONstruct := Profile{
		Email:     respSQLstruct.Email,
		Phone:     respSQLstruct.Phone,
		Roles:     respSQLstruct.Roles,
		FirstName: respSQLstruct.FirstName,
		LastName:  respSQLstruct.LastName,
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(respJSONstruct)
	if err != nil {
		slog.Error("failed writing json to response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// UpdateProfileByID implements ServerInterface.
func (s Server) UpdateProfileByID(w http.ResponseWriter, r *http.Request, profileID string) {
	dataJSONstruct := Profile{}
	err := json.NewDecoder(r.Body).Decode(&dataJSONstruct)
	if err != nil {
		slog.Error("failed reading json request", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataSQLstruct := models.Profile{
		FirstName: dataJSONstruct.FirstName,
		LastName:  dataJSONstruct.LastName,
		Email:     dataJSONstruct.Email,
		Phone:     dataJSONstruct.Phone,
		Roles:     dataJSONstruct.Roles,
	}

	err = models.UpdateProfile(s.db, dataSQLstruct, profileID)
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

// GetAPIToken implements ServerInterface.
func (s Server) GetAPIToken(w http.ResponseWriter, r *http.Request) {
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

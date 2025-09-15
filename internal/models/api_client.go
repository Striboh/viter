// Package models provides instances of entries of DB
package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Profile represents a user's profile information and is managed by sqlx to retrieve info
type ApiClient struct {
	ID        uuid.UUID `db:"id"`
	Hash      string
	Roles     []string
}

// GetProfile retrieves Profile from db
func GetApiClient(db *sqlx.DB, id uuid.UUID) (ApiClient, error) {
	var entry ApiClient
	err := db.Get(&entry, "SELECT * FROM api_clients WHERE id = '$1'", id)
	if err != nil {
		return ApiClient{}, fmt.Errorf("failed selecting profile with id %s: %w", id, err)
	}
	entry.Roles, err = GetRoles(db, id)
	if err != nil {
		return ApiClient{}, err
	}
	return entry, nil
}

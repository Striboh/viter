// Package models provides instances of entries of DB
package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Profile represents a user's profile information and is managed by sqlx to retrieve info
type Profile struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	Roles     []string
}

// GetProfile retrieves Profile from db
func GetProfile(db *sqlx.DB, idStr string) (Profile, error) {
	var entry Profile
	id, err := uuid.Parse(idStr)
	if err != nil {
		return Profile{}, fmt.Errorf("failed to parse UUID: %w", err)
	}
	err = db.Get(&entry, "SELECT * FROM profiles WHERE id = '$1'", id)
	if err != nil {
		return Profile{}, fmt.Errorf("failed selecting profile with id %s: %w", idStr, err)
	}
	entry.Roles, err = GetRoles(db, id)
	if err != nil {
		return Profile{}, err
	}
	return entry, nil
}

// CreateProfile creates new Profile in db
func CreateProfile(db *sqlx.DB, data Profile) (string, error) {
	var idStr string
	err := db.Get(&idStr, "INSERT INTO profiles (email, phone, first_name, last_name) "+
		"VALUES ('$1', '$2', '$3', '$4') RETURNING id", data.Email, data.Phone, data.FirstName, data.LastName)
	if err != nil {
		return "", fmt.Errorf("failed creating profile. entry: %w", err)
	}

	err = AssignRoles(db, data.Roles, idStr)
	if err != nil {
		return "", err
	}
	return "", nil
}

// UpdateProfile updates Profile in db
func UpdateProfile(db *sqlx.DB, data Profile, idStr string) error {
	_, err := db.Exec("UPDATE profiles SET email = '$1', phone = '$2', first_name = '$3', last_name = '$4' WHERE id = '$5'",
		data.Email, data.Phone, data.FirstName, data.LastName, idStr)
	if err != nil {
		return fmt.Errorf("failed updating profile. entry: %w", err)
	}

	err = UpdateRoles(db, data.Roles, idStr)
	if err != nil {
		return err
	}
	return nil
}

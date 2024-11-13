// Package models provides instances of entries of DB
package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// GetRoles retrieves roles list for given Profile
func GetRoles(db *sqlx.DB, profileID uuid.UUID) ([]string, error) {
	var roles []string

	err := db.Select(&roles, "SELECT role FROM role_records WHERE profile_id = '$1'", profileID)
	if err != nil {
		return nil, fmt.Errorf("failed selecting roles for profile with id %s. entry: %w", profileID.String(), err)
	}

	return roles, nil
}

// AssignRoles assigns roles for newly created Profile
func AssignRoles(db *sqlx.DB, roles []string, profileIDStr string) error {
	queryStr := "INSERT INTO role_records (profile_id, role) VALUES"
	for i := range len(roles) - 1 {
		queryStr += fmt.Sprintf("\n('%s', '%s'),", profileIDStr, roles[i])
	}
	queryStr += fmt.Sprintf("\n('%s', '%s');", profileIDStr, roles[len(roles)-1])

	_, err := db.Exec(queryStr)

	if err != nil {
		return fmt.Errorf("failed assigning roles for profile with id %s. entry: %w", profileIDStr, err)
	}
	return nil
}

// UpdateRoles updates list of roles for given Profile
func UpdateRoles(db *sqlx.DB, roles []string, profileIDStr string) error {
	queryStr := "DELETE FROM role_records WHERE profile_id = '" + profileIDStr + "';" +
		"\nINSERT INTO role_records (profile_id, role) VALUES"
	for i := range len(roles) - 1 {
		queryStr += fmt.Sprintf("\n('%s', '%s'),", profileIDStr, roles[i])
	}
	queryStr += fmt.Sprintf("\n('%s', '%s');", profileIDStr, roles[len(roles)-1])

	_, err := db.Exec(queryStr)

	if err != nil {
		return fmt.Errorf("failed assigning roles for profile with id %s. entry: %w", profileIDStr, err)
	}
	return nil
}

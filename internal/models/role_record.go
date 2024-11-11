// Package models provides instances of entries of DB
package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func GetRoles(db *sqlx.DB, ProfileID uuid.UUID) ([]string, error) {
	var roles []string

	err := db.Select(&roles, "SELECT role FROM role_records WHERE profile_id = '$1'", ProfileID)
	if err != nil {
		return nil, fmt.Errorf("failed selecting roles for profile with id %s. entry: %w", ProfileID.String(), err)
	}

	return roles, nil
}
func AssignRoles(db *sqlx.DB, roles []string, ProfileID_str string) error {
	query_str := "INSERT INTO role_records (profile_id, role) VALUES"
	for i := range len(roles) - 1 {
		query_str += fmt.Sprintf("\n('%s', '%s'),", ProfileID_str, roles[i])
	}
	query_str += fmt.Sprintf("\n('%s', '%s');", ProfileID_str, roles[len(roles)-1])

	_, err := db.Exec(query_str)

	if err != nil {
		return fmt.Errorf("failed assigning roles for profile with id %s. entry: %w", ProfileID_str, err)
	}
	return nil
}

func UpdateRoles(db *sqlx.DB, roles []string, ProfileID_str string) error {
	query_str := "DELETE FROM role_records WHERE profile_id = '" + ProfileID_str + "';" +
		"\nINSERT INTO role_records (profile_id, role) VALUES"
	for i := range len(roles) - 1 {
		query_str += fmt.Sprintf("\n('%s', '%s'),", ProfileID_str, roles[i])
	}
	query_str += fmt.Sprintf("\n('%s', '%s');", ProfileID_str, roles[len(roles)-1])

	_, err := db.Exec(query_str)

	if err != nil {
		return fmt.Errorf("failed assigning roles for profile with id %s. entry: %w", ProfileID_str, err)
	}
	return nil
}

// Package models provides instances of entries of DB
package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// KeyValue - entry of `key_value` table
type KeyValue struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

// KeyValueFromDB - retrieved KeyValue from database.
func KeyValueFromDB(db *sqlx.DB, key string) (KeyValue, error) {
	var entry KeyValue
	
	err := db.Get(&entry, "SELECT * FROM key_value WHERE key = $1", key)
	if err != nil {
		return KeyValue{}, fmt.Errorf("failed selecting key_value with key %s entry: %w", key, err)
	}

	return entry, nil
}

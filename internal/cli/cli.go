// Package cli provides parsed command line parameters.
package cli

import (
	"flag"
	"log"
	"slices"
)

// MigrateMode - enum of possible modes of migration set by user
type MigrateMode int

const (
	// MigrateModeUp - apply newest migrations
	MigrateModeUp MigrateMode = iota + 1
	// MigrateModeDown - revert migrations
	MigrateModeDown
	// MigrateModeSkip - skip migrations
	MigrateModeSkip
)

var (
	migrateMode *string
)

func init() {
	migrateMode = flag.String("migrate", "", "Specify 'up' or 'down'")
	flag.Bool("h", false, "use to print help")
	flag.Parse()

	if !slices.Contains([]string{"down", "up", ""}, *migrateMode) {
		log.Fatalf("migrate mode should be 'up', 'down' or nothing, got %s", *migrateMode)
	}
}

// GetMigrateMode - return migration mode specified by user in command
// line arguments.
func GetMigrateMode() MigrateMode {
	switch *migrateMode {
	case "up":
		return MigrateModeUp
	case "down":
		return MigrateModeDown
	default:
		return MigrateModeSkip
	}
}

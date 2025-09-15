// Package config for configuration options of the service
package config

import (
	"cmp"
	"os"
)

const (
	portEnv        = "VITER_PORT"
	hostEnv        = "VITER_HOST"
	postgresURLEnv = "VITER_POSTGRES"
	jwtSecretEnv   = "VITER_JWT_SECRET"

	portDefault = "8080"
	hostDefault = "127.0.0.1"
	//nolint:gosec
	postgresURLDefault = "postgresql://viter:admin123@127.0.0.1:15432/viterdb?sslmode=disable"

	jwtSecret = "secret"
	
)

// Config - configurable options of the service
type Config struct {
	ListenAddr  string
	PostgresURL string
	JwtSecret   string
}

// FromEnv - creates Config from environmental variables.
func FromEnv() Config {
	config := Config{
		ListenAddr:  hostDefault + ":" + portDefault,
		PostgresURL: postgresURLDefault,
		JwtSecret: jwtSecret,
	}

	port := os.Getenv(portEnv)
	host := os.Getenv(hostEnv)

	if port != "" || host != "" {
		config.ListenAddr = cmp.Or(host, hostDefault) + ":" + cmp.Or(port, portDefault)
	}

	if val := os.Getenv(postgresURLEnv); val != "" {
		config.PostgresURL = val
	}
	if val := os.Getenv(jwtSecretEnv); val != "" {
		config.JwtSecret = val
	}

	return config
}

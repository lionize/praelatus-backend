package config

import "os"

// GetDbURL will return the environment variable PRAELATUS_DB if set, otherwise
// return the default development database url.
func GetDbURL() string {
	url := os.Getenv("PRAELATUS_DB")
	if url == "" {
		return "postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable"
	}

	return url
}

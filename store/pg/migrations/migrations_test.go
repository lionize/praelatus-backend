package migrations

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func testDB(t *testing.T) *sql.DB {
	d, err := sql.Open("postgres",
		"postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	return d
}

func TestRunMigrations(t *testing.T) {
	db := testDB(t)

	err := RunMigrations(db)
	if err != nil {
		t.Error(err)
	}
}

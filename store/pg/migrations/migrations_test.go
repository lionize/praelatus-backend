package migrations

import (
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
)

type TestSQLStore struct {
	db *sqlx.DB
}

func (t *TestSQLStore) SchemaVersion() int {
	return 0
}

func (t *TestSQLStore) RunQuery(query string) (*sqlx.Rows, error) {
	return t.db.Queryx(query)
}

func (t *TestSQLStore) RunExec(query string) (sql.Result, error) {
	return t.db.Exec(query)
}

func TestRunMigrations(t *testing.T) {

}

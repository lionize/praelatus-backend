package pg

import "github.com/jmoiron/sqlx"

type WorkflowStore struct {
	db *sqlx.DB
}

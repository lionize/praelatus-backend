package pg

import "github.com/jmoiron/sqlx"

// WorkflowStore contains methods for saving/retrieving workflows from a
// postgres DB
type WorkflowStore struct {
	db *sqlx.DB
}

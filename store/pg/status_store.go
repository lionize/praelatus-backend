package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// StatusStore contains methods for storing and retrieving Statuses from a
// Postgres DB
type StatusStore struct {
	db *sqlx.DB
}

// Get TODO
func (ss *StatusStore) Get(ID int) (*models.Status, error) {
	return nil, nil
}

// New TODO
func (ss *StatusStore) New(status *models.Status) error {
	return nil
}

// Save TODO
func (ss *StatusStore) Save(status *models.Status) error {
	return nil
}

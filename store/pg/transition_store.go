package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// TransitionStore contains methods for storing and retrieving Transitions from
// a Postgres DB
type TransitionStore struct {
	db *sqlx.DB
}

// Get TODO
func (ts *TransitionStore) Get(ID int) (*models.Transition, error) {
	return nil, nil
}

// New TODO
func (ts *TransitionStore) New(transition *models.Transition) error {
	return nil
}

// Save TODO
func (ts *TransitionStore) Save(transition *models.Transition) error {
	return nil
}

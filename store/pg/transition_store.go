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

// Get gets a transition by it's ID from a postgres DB.
func (ts *TransitionStore) Get(ID int) (*models.Transition, error) {
	var s models.Transition
	err := ts.db.QueryRowx("SELECT * FROM transitions WHERE id = $1", ID).
		StructScan(&s)
	return &s, err
}

// New will create a new Transition in the postgres DB.
func (ts *TransitionStore) New(transition *models.Transition) error {
	id, err := ts.db.Exec(`INSERT INTO transitions VALUES
	(name, workflow_id, status_id) = (?, ?, ?)`,
		transition.Name, transition.WorkflowID, transition.StatusID)
	if err != nil {
		return err
	}

	transition.ID, err = id.LastInsertId()
	return err
}

// Save TODO
func (ts *TransitionStore) Save(transition *models.Transition) error {
	return nil
}

package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

// TransitionStore contains methods for storing and retrieving Transitions from
// a Postgres DB
type TransitionStore struct {
	db *sql.DB
}

// Get gets a transition by it's ID from a postgres DB.
func (ts *TransitionStore) Get(ID int64) (*models.Transition, error) {
	var s models.Transition
	err := ts.db.QueryRowx("SELECT * FROM transitions WHERE id = $1", ID).
		StructScan(&s)
	return &s, handlePqErr(err)
}

// New will create a new Transition in the postgres DB.
func (ts *TransitionStore) New(transition *models.Transition) error {
	err := ts.db.QueryRow(`INSERT INTO transitions 
						   (name, workflow_id, status_id) 
						   VALUES ($1, $2, $3)
						   RETURNING id;`,
		transition.Name, transition.WorkflowID, transition.StatusID).
		Scan(&transition.ID)

	return handlePqErr(err)
}

// Save update an existing Transition in the postgres DB.
func (ts *TransitionStore) Save(transition *models.Transition) error {
	_, err := ts.db.Exec(`UPDATE transitions SET
						  (name, workflow_id, status_id) = ($1, $2, $3)
						  WHERE id = $4`,
		transition.Name, transition.WorkflowID,
		transition.StatusID, transition.ID)
	return handlePqErr(err)
}

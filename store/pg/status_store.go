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

// Get gets a Status by it's ID in a postgres DB
func (ss *StatusStore) Get(ID int64) (*models.Status, error) {
	var s models.Status
	err := ss.db.QueryRowx("SELECT * FROM statuses WHERE id = $1;", ID).
		StructScan(&s)
	return &s, err
}

// New creates a new Status in the postgres DB
func (ss *StatusStore) New(status *models.Status) error {
	err := ss.db.QueryRow(`INSERT INTO statuses (name) VALUES ($1)
						   RETURNING id;`,
		status.Name).
		Scan(&status.ID)

	return handlePqErr(err)
}

// Save updates a Status in the postgres DB
func (ss *StatusStore) Save(status *models.Status) error {
	_, err := ss.db.Exec(`UPDATE statuses SET (name) = ($1);`, status.Name)
	return err
}
